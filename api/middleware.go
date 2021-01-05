package api

import (
	"bufio"
	"errors"
	"net"
	"net/http"
)

//	Wrap handler in own structure to redirect to an own 404-Handler.
type handlerSet struct {
	http.Handler
	redirect404 bool
}

type Middleware []handlerSet

type MiddlewareResponseWriter struct {
	http.ResponseWriter
	isWritten   bool
	redirect404 bool
	forbidWrite bool
	hijacker    http.Hijacker
}

func (m *Middleware) Add(handler http.Handler, redirect404 bool) {
	newHandler := handlerSet{handler, redirect404}
	*m = append(*m, newHandler)
}

func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mw := NewMiddlewareResponseWriter(w)
	origin := r.Header.Get("Origin")
	if origin == "" {
		origin = "*"
	}
	mw.Header().Set("Access-Control-Allow-Origin", origin)
	mw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	mw.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	for _, handler := range m {
		mw.SetRedirect404(handler.redirect404)
		handler.ServeHTTP(mw, r)
		if mw.isWritten {
			return
		}
	}
	Handle404(w, r)
}

func NewMiddlewareResponseWriter(w http.ResponseWriter) *MiddlewareResponseWriter {
	hijacker, _ := w.(http.Hijacker)
	return &MiddlewareResponseWriter{
		ResponseWriter: w,
		hijacker:       hijacker,
	}
}

func (w *MiddlewareResponseWriter) Write(bytes []byte) (int, error) {
	if w.forbidWrite {
		return 0, nil
	}
	w.isWritten = true
	return w.ResponseWriter.Write(bytes)
}

func (w *MiddlewareResponseWriter) WriteHeader(code int) {
	if w.redirect404 && code == 404 {
		Handle404(w, nil)
		w.forbidWrite = true
		return
	}
	w.isWritten = true
	w.ResponseWriter.WriteHeader(code)
}

//	If Redirect404 is set to true, the middleware will serve Handle404() whenever
//	a handler tries to write 404.
func (w *MiddlewareResponseWriter) SetRedirect404(value bool) {
	w.redirect404 = value
}

//  Make the ResponseWriter support Hijacker in order to enable Websockets
func (r *MiddlewareResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if r.hijacker == nil {
		return nil, nil, errors.New("http.Hijacker not implemented by underlying http.ResponseWriter")
	}
	c, b, e := r.hijacker.Hijack()
	if e == nil {
		r.forbidWrite = true
	}
	return c, b, e
}

package main

import (
	"github.com/gnampfelix/grender/api"
	"github.com/gnampfelix/grender/geometry"
	"github.com/gnampfelix/grender/renderer"
	"github.com/gnampfelix/pub"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type NotFound struct{}

func (n *NotFound) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	notFound := new(NotFound)
	router.NotFound = notFound
	return router
}

func main() {
	myPub = pub.New()
	api.SetPublisher(myPub)

	object := renderer.NewCube()
	input := renderer.NewInput()
	input.Add(object)
	output := NewStreamerOutput(270, 480)
	renderer := renderer.New()

	go func() {
		for {
			object.Rotate(geometry.Z, 1.2)
			renderer.Render(input, output)
		}
	}()

	frontendRouter := NewRouter()
	frontendRouter.ServeFiles("/*filepath", http.Dir("html/"))

	router := NewRouter()
	router.GET("/api/ws", api.ServeWebsocket)

	middleware := api.Middleware{}
	middleware.Add(router, false)        //--> The api could return a 404 in terms of "ressource not found"
	middleware.Add(frontendRouter, true) //--> Don't handle 404-Erros, let the frontend handle them!
	log.Println("Starting server..")
	log.Fatal(http.ListenAndServe(":8000", middleware))
}

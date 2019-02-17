package main

import (
	"github.com/julienschmidt/httprouter"
  "github.com/gnampfelix/pub"
	"felix/canvas/api"
  "log"
  "net/http"
	"github.com/gnampfelix/grender/geometry"
	. "github.com/gnampfelix/grender/renderer"
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

	object := NewCube()
	input := NewInput()
	input.Add(object)
	renderer := New(NewStreamerOutput(270, 480))

	for i := 0; i < 1; i++ {
		object.Rotate(geometry.Z, 1.2)
		go renderer.Render(input)
}

  frontendRouter := NewRouter()
  frontendRouter.ServeFiles("/*filepath", http.Dir("html/"))

	router := NewRouter()
  router.GET("/api/ws", api.ServeWebsocket)

  middleware := api.Middleware{}
  middleware.Add(router, false) //--> The api could return a 404 in terms of "ressource not found"
  middleware.Add(frontendRouter, true) //--> Don't handle 404-Erros, let the frontend handle them!
	log.Println("Starting server..")
  log.Fatal(http.ListenAndServe(":8000", middleware))
}

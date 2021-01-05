package api

import (
	"github.com/gnampfelix/pub"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
)

var upgrader websocket.Upgrader
var conn []*websocket.Conn
var publisher pub.Publisher

func init() {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn = make([]*websocket.Conn, 0)
}

func SetPublisher(p pub.Publisher) {
	publisher = p
	go writeToSocket()
}

func ServeWebsocket(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ws, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Serving Websocket..")
	conn = append(conn, ws)
}

func writeToSocket() {
	sub := publisher.Subscribe("ws", pub.NewSubscriber)
	for {
		message := sub.WaitForMessage()
		content, err := ioutil.ReadAll(message)
		if err != nil {
			log.Println(err)
			continue
		}
		for i, c := range conn {
			err = c.WriteMessage(websocket.TextMessage, content)
			if err != nil {
				log.Println(err)
				c.Close()
				if i+1 >= len(conn) {
					conn = conn[:i]
				} else {
					conn = append(conn[:i], conn[i+1:]...)
				}
			}
		}
	}
}

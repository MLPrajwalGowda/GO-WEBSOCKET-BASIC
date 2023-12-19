package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Connecting to websocket named Userwebsocket func
func UserWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		// read message from client
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		// show message
		log.Printf("Message Received : %s", message)

		// send message to client
		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println(err)
			break
		}

	}
}

func main() {
	http.HandleFunc("/socket", UserWebsocket)
	fmt.Println("connect to server -- ws://localhost:8080 ")
	// ws://localhost:8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

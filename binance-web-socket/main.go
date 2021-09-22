package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	//"time"
)

var upgrader = websocket.Upgrader{}

var messageSent = []byte("Pong")
func handler(writer http.ResponseWriter, request *http.Request) {
	var conn, _ = upgrader.Upgrade(writer, request, nil)
	go func(conn *websocket.Conn) {
		for {
			_, messageRecieved, _ := conn.ReadMessage()
			fmt.Println(string(messageRecieved))
			conn.WriteMessage(1, messageSent)
			fmt.Println(string(messageSent))
		}
	}(conn)
}

func main() {
	http.HandleFunc("/ws", handler)

	http.ListenAndServe(":3000", nil)
}

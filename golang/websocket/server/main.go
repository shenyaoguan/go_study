package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var conns []*websocket.Conn

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := UP.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	conns = append(conns, conn)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		for i := range conns {
			conns[i].WriteMessage(websocket.TextMessage, []byte("你说的是："+string(msg)+"吗？"))
		}
		defer conn.Close()
		log.Println(string(msg))
		log.Println("Message sent")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

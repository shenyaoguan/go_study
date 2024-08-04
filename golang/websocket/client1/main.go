package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
)

func main() {
	dl := websocket.Dialer{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, _, err := dl.Dial("ws://localhost:8080", nil)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	conn.WriteMessage(websocket.TextMessage, []byte("Hello, world!"))
	go send(conn)
	for {
		m, p, e := conn.ReadMessage()
		if e != nil {
			log.Println(e)
			panic(e)
		}
		fmt.Println(m, string(p))
	}
}

func send(conn *websocket.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		l, _, _ := reader.ReadLine()
		conn.WriteMessage(websocket.TextMessage, l)
	}
}

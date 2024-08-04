package main

import (
	"fmt"
	"net"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:8080")
	TCPListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer TCPListener.Close()
	for {
		conn, err := TCPListener.AcceptTCP()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			return
		}
		go handleConnetion(conn)
	}

}

func handleConnetion(conn *net.TCPConn) {
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read from conn failed, err:", err)
			break
		}
		fmt.Println(conn.RemoteAddr().String() + " " + string(buf[0:n]) + " connected")
		str := "收到了: " + string(buf[0:n])
		conn.Write([]byte(str))
	}

}

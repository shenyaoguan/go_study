package main

import (
	"bufio"
	"net"
	"os"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:8080")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	reader := bufio.NewReader(os.Stdin)
	for {
		bytes, _, _ := reader.ReadLine()
		conn.Write(bytes)
		rb := make([]byte, 1024)
		rn, _ := conn.Read(rb)
		println(string(rb[0:rn]))
	}
}

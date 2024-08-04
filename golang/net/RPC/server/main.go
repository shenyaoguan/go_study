package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Server struct {
}

type Req struct {
	NumOne int
	NumTwo int
}

type Res struct {
	Result int
}

func (s Server) GetName(req Req, res *Res) error {
	time.Sleep(5 * time.Second)
	res.Result = req.NumOne + req.NumTwo
	return nil

}

func main() {
	rpc.Register(new(Server))
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
	}
	http.Serve(l, nil)
}

package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Req struct {
	NumOne int
	NumTwo int
}

type Res struct {
	Result int
}

func main() {
	req := Req{NumOne: 1, NumTwo: 2}
	var res Res
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	ca := client.Go("Server.GetName", req, &res, nil)

	for {
		select {
		case <-ca.Done:
			fmt.Println(res)
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("waiting")
		}
	}

}

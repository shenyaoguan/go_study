package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	message := make(chan int)
	go son(ctx, message)
	for i := 0; i < 10; i++ {
		message <- i
	}
	defer cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("结束")
}

func son(ctx context.Context, msg chan int) {
	t := time.Tick(1 * time.Second)
	for _ = range t {
		select {
		case msg := <-msg:
			fmt.Println("msg:", msg)
		case <-ctx.Done():
			fmt.Println("end")
			return
		}
	}
	return
}

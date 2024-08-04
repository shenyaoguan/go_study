package main

import (
	"fmt"
	"sync"
	"time"
)

func SyncClass() {
	l := &sync.Mutex{}
	go lockFun(l)
	go lockFun(l)
	go lockFun(l)
	go lockFun(l)
}

func lockFun(lock *sync.Mutex) {
	lock.Lock()
	fmt.Println("lockFun")
	time.Sleep(1 * time.Second)
	lock.Unlock()
}

func main() {
	SyncClass()
	time.Sleep(1 * time.Minute)
}

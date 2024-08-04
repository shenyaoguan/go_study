package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("close file failed, err:", err)
			return
		}
	}(f)
	reader := bufio.NewReader(f)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Println(str)
	}
}

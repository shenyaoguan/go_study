package main

import (
	"bytes"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("POST", "http://localhost:8080/test", bytes.NewBuffer([]byte("{\"test\": \"test\"}")))
	request.Header["test"] = []string{"test1", "test2"}
	if err != nil {
		panic(err)
	}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	b, _ := io.ReadAll(response.Body)
	println(string(b))
	defer response.Body.Close()
}

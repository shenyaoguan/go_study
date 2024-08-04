package main

import (
	"io"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		res.Write([]byte("GET"))
		break
	case "POST":
		b, _ := io.ReadAll(req.Body)
		header := res.Header()
		header["test"] = []string{"test1", "test2"}
		res.WriteHeader(http.StatusBadRequest)
		res.Write(b)
		break
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", handler)
	http.ListenAndServe(":8080", mux)
}

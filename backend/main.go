package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("test"))
	})
	http.ListenAndServe(":8080", nil)
}

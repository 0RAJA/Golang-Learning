package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct{}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World-->详细配置")
}

func main() {
	myHandler := MyHandler{}
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler,
		TLSConfig:   nil,
		ReadTimeout: 2 * time.Second,
	}
	server.ListenAndServe()
}

package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	//只要实现了ServeHTTP方法就可以作为处理器
	myHandler := MyHandler{}
	http.Handle("/go", &myHandler)
	http.ListenAndServe(":8080", nil)
}

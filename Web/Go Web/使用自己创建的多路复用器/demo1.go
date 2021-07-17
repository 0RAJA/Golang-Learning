package main

import (
	"fmt"
	"net/http"
)

type Person struct {
	name string
}
func handConn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello-->mux")
}

func main() {
	mux := http.NewServeMux() //创建自己的多路复用器

	mux.HandleFunc("/mux", handConn) //注册处理器

	http.ListenAndServe(":8080", mux) //监听端口
}

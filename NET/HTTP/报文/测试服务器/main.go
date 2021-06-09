package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

//测试服务器
func main() {
	http.HandleFunc("/go", myHandler)

	//在指定地址监听,开启一个HTTP
	http.ListenAndServe("127.0.0.1:8000", nil)
}

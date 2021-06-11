package main

import (
	"fmt"
	"net/http"
)

// HandConn 创建处理器函数
func HandConn(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method) // Method指定HTTP方法（GET、POST、PUT等）。对客户端，""代表GET。
	fmt.Println(req.URL)    // URL在服务端表示被请求的URI，在客户端表示要访问的URL
	fmt.Println(req.Header) //Header字段用来表示HTTP请求的头域。
	fmt.Println(req.Body)   //Body是请求的主体。
	w.Write([]byte("hello go"))
}

func main() {
	//注册处理函数,用户连接,自动调用指定的处理函数
	http.HandleFunc("/go", HandConn)

	//创建路由
	http.ListenAndServe(":8000", nil)
}

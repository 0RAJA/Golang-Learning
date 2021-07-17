package main

import (
	"fmt"
	"net/http"
)

/*
Go提供了一系列用于创建Web服务器的标准库,只要通过net/http包调用ListenAndServe函数并传入网络地址以及负
责处理请求的处理器( handler )作为参数就可以了。如果网络地址参数为空字符串，那
么服务器默认使用80端口进行网络连接;如果处理器参数为nil,那么服务器将使用默
认的多路复用器DefaultServeMux,当然,我们也可以通过调用NewServeMux函数创
建一个多路复用器。多路复用器接收到用户的请求之后根据请求的URL来判断使用哪
个处理器来处理请求，找到后就会重定向到对应的处理器来处理请求,
*/
//创建处理器函数
func handConn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world\n"))
	fmt.Fprintln(w, "NIHAO", r.URL.Path)
}

func main() {
	//注册处理器-->把普通函数转换为处理器函数
	http.HandleFunc("/gogo", handConn)
	//创建路由
	http.ListenAndServe(":8080", nil)
}

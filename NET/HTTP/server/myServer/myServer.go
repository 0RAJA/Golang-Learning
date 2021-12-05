package main

import (
	"net/http"
	"time"
)

type MyHandler struct{}

//自己的处理器
func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("自己的详细处理器"))
}

func main() {
	myHandler := MyHandler{}
	//注册处理器
	//http.Handle("/go", &myHandler)
	//设置服务器参数
	server := http.Server{
		Addr:      ":8080", // 监听的TCP地址，如果为空字符串会使用":http"
		TLSConfig: nil,     //可选的TLS配置，用于ListenAndServeTLS方法
		//ReadTimeout: 2 * time.Second, // 请求的读取操作在超时前的最大持续时间
		Handler: http.TimeoutHandler(&myHandler, time.Second*2, "服务端操作超时"),
	}
	server.ListenAndServe()
	//http.ListenAndServe(":8080", nil)
}

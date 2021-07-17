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
//GET
///go
//map[Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9] Accept-Encoding:[gzip, deflate, br] Accept-Language:[zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6] Connection:[keep-alive] Dnt:[1] Sec-Ch-Ua:["Chromium";v="92", " Not A;Brand";v="99", "Microsoft Edge";v="92"] Sec-Ch-Ua-Mobile:[?0] Sec-Fetch-Dest:[document] Sec-Fetch-Mode:[navigate] Sec-Fetch-Site:[none] Sec-Fetch-User:[?1] Upgrade-Insecure-Requests:[1] User-Agent:[Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.93 Safari/537.36 Edg/92.0.902.45]]
//{}
package main

import (
	"fmt"
	"net"
)

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	requestBuf := `GET / HTTP/1.1		
Host: 127.0.0.1:8000
Connection: keep-alive
sec-ch-ua: " Not;A Brand";v="99", "Microsoft Edge";v="91", "Chromium";v="91"
sec-ch-ua-mobile: ?0
DNT: 1
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36 Edg/91.0.864.37
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*; q = 0.8, application/signed-exchange; v = b3; q= 0.9
Sec-Fetch-Site: none
Sec-Fetch-Mode: navigate
Sec-Fetch-User: ?1
Sec-Fetch-Dest: document
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN, zh; q = 0.9, en; q= 0.8
sec-gpc: 1
					   
					   
`
	//发送请求包,服务器才回复响应包
	_, _ = conn.Write([]byte(requestBuf))

	//接受响应包
	buf := make([]byte, 1024*4)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	//打印响应报文
	fmt.Printf("#\n%v\n#", string(buf[:n]))
}

//响应报文:
//HTTP/1.1 400 Bad Request		//相应行
//Content-Type: text/plain; charset=utf-8	//响应头
//Connection: close
//								//空行
//400 Bad Request				//响应包体

package main

import (
	"fmt"
	"net"
)

func main() {
	lister, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		return
	}
	defer lister.Close()
	conn, err := lister.Accept()
	if err != nil {
		return
	}
	defer conn.Close()
	buf := make([]byte, 1024*4)
	n, err := conn.Read(buf)
	if err != nil {
		return
	}
	fmt.Printf("#\n%v\n#", string(buf[:n]))
}

//请求报文
//GET / HTTP/1.1		//请求行GET->
//Host: 127.0.0.1:8000	//请求头
//Connection: keep-alive
//sec-ch-ua: " Not;A Brand";v="99", "Microsoft Edge";v="91", "Chromium";v="91"
//sec-ch-ua-mobile: ?0
//DNT: 1
//Upgrade-Insecure-Requests: 1
//User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36 Edg/91.0.864.37
//Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*; q = 0.8, application/signed-exchange; v = b3; q= 0.9
//Sec-Fetch-Site: none
//Sec-Fetch-Mode: navigate
//Sec-Fetch-User: ?1
//Sec-Fetch-Dest: document
//Accept-Encoding: gzip, deflate, br
//Accept-Language: zh-CN, zh; q = 0.9, en; q= 0.8
//sec-gpc: 1
//					   //空行
//					   //请求包体

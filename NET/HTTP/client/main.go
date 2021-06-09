package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8000/go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()      //结束时记住关闭包体
	fmt.Println(resp.Status)     //状态
	fmt.Println(resp.StatusCode) //状态码
	fmt.Println(resp.Header)     //头信息
	fmt.Println("++++++++++++++++++++++++++++")
	//resp.Body //流
	buf := make([]byte, 4*1024)
	var msg string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println(err)
			break
		}
		msg += string(buf[:n])
	}
	fmt.Println(msg)
}

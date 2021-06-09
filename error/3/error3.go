package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.LookupHost("www.google.com")
	if err != nil {
		if ins, ok := err.(*net.DNSError); ok {
			if ins.Timeout() {
				fmt.Println("超时")
			} else if ins.Temporary() {
				fmt.Println("临时性错误")
			} else {
				fmt.Println("通常错误")
			}
		}
		return
	}
	fmt.Println(addr)
}

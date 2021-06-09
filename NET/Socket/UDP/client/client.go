package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.Dial("udp", "127.0.0.1:8800")
	if err != nil {
		return
	}
	defer socket.Close()
	_, err = socket.Write([]byte("你好"))
	if err != nil {
		return
	}
	buf := make([]byte, 1024)
	n, err := socket.Read(buf)
	fmt.Println(string(buf[:n]))

}

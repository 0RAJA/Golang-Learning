package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	_, err := conn.Write([]byte("Hello world"))
	if err != nil {
		return
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return
	}
	fmt.Println(string(buf[:n]))
}

func main() {
	lister, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8800,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer lister.Close()
	buf := make([]byte, 1024)
	for {
		n, addr, err := lister.ReadFromUDP(buf)
		if err != nil {
			return
		}
		fmt.Println(string(buf[:n]))
		_, err = lister.WriteToUDP([]byte("hello"), addr)
		if err != nil {
			return
		}
	}

}

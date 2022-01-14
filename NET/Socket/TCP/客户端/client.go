package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
)

// PrintLocalAddr 展示本机地址
func PrintLocalAddr(conn net.Conn) {
	fmt.Println("您的本机地址为:", conn.LocalAddr().String())
}

// IllegalInput 判断输入的字符串是否合法
func IllegalInput(s string) bool {
	if strings.Contains(s, "|") {
		return false
	}
	return true
}

func SendAccount(conn net.Conn, inputReader *bufio.Reader) bool {
	var options int
	for {
		_ = exec.Command("CLS").Run() //清屏
		fmt.Print("1.登录账号\n2.注册账号\n3.退出\n")
		_, err := fmt.Fscanf(inputReader, "%d\n", &options) //读入选项
		if err != nil || options < 1 || options > 3 {
			fmt.Println("输入有误请重新输入")
			continue
		}
		var (
			name     string
			password string
		)
		for {
			switch options {
			case 1:
				fmt.Print("输入姓名:")
				name, _ = inputReader.ReadString('\n')
				fmt.Print("输入密码:")
				password, _ = inputReader.ReadString('\n')
				name, password = strings.Trim(name, "\n"), strings.Trim(password, "\n")
				if IllegalInput(name) == false || IllegalInput(password) == false {
					fmt.Println("姓名或密码存在非法字符'|'")
					break
				}
				_, err := conn.Write([]byte(name + "|" + password))
				if err != nil {
					fmt.Println("登录信息发送失败,err = ", err)
					return false
				}
				test := make([]byte, 1024) //读取服务器返回的验证信息
				n, err := conn.Read(test)
				if err != nil {
					fmt.Println("返回验证信息失败,err = ", err)
					return false
				}
				if string(test[:n]) == "OK" {
					fmt.Println("登陆成功")
					return true
				} else {
					fmt.Println("用户名或密码错误")
					return false
				}
			case 2:
				return false
			case 3:
				return false
			}
		}
	}
}

func SendMessage(conn net.Conn, inputReader *bufio.Reader) {
	fmt.Println("输入信息,输入q退出")
	for {
		input, _ := inputReader.ReadString('\n')
		input = strings.Trim(input, "\r\n")
		if strings.ToUpper(input) == "Q" {
			fmt.Println("Bye")
			return
		}
		_, err := conn.Write([]byte(input))
		if err != nil {
			fmt.Println("send error,err:", err)
			continue
		}
		reply := make([]byte, 1024)
		_, err = conn.Read(reply) //接受返回信息
		if err != nil {
			fmt.Println("reply error,err:", err)
		}
	}
}

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:80") //拨号连接
	if err != nil {
		fmt.Println("err == ", err)
		return
	}
	defer conn.Close()
	PrintLocalAddr(conn)
	//发送数据
	inputReader := bufio.NewReaderSize(os.Stdin, 1024)
	if SendAccount(conn, inputReader) {
		SendMessage(conn, inputReader)
	}
}

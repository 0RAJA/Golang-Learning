package main

import (
	"go-gomail/gomail"
	"strconv"
)

func SendMail(mailTo string, subject string, body string) error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": "1647193241@qq.com",
		"pass": "ytulmprmrlpmceah",
		"host": "smtp.qq.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From", "XD Game"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo)                               //发送给多个用户
	m.SetHeader("Subject", subject)                         //设置邮件主题
	m.SetBody("textml", body)                               //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}
func main() {
	//定义收件人
	mailTo := "2682072219@qq.com"
	//邮件主题为"Hello"
	subject := "gomail测试你是不是傻狗"
	// 邮件正文
	body := "测试完毕,你是傻狗"
	SendMail(mailTo, subject, body)
}

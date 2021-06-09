package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net"
	"strings"
)

const (
	driverName = "mysql"                            //驱动名-这个名字其实就是数据库驱动注册到database/sql 时所使用的名字.
	userName   = "root"                             //用户名
	passWord   = "WW876001"                         //密码
	ip         = "127.0.0.1"                        //ip地址
	port       = "3306"                             //端口号
	dbName     = "tell"                             //数据库名
	ipFormat   = "%s:%s@tcp(%s:%s)/%s?charset=utf8" //格式
)

type Person struct {
	name     string
	password string
}

// DBInit 初始化数据库
func DBInit() (*sql.DB, error) {
	DB, err := sql.Open(driverName, fmt.Sprintf(ipFormat, userName, passWord, ip, port, dbName))
	if err != nil {
		return nil, err
	}
	DB.SetMaxOpenConns(100) //SetMaxOpenConns设置与数据库建立连接的最大数目。
	DB.SetMaxIdleConns(10)  //SetMaxIdleConns设置连接池中的最大闲置连接数。
	err = DB.Ping()
	if err != nil {
		return nil, err
	}
	return DB, nil
}

// LegalAccount 验证账户合法性
func LegalAccount(conn net.Conn) (string, bool) {
	logIn := make([]byte, 1024)
	n, err := conn.Read(logIn)
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	message := strings.Split(string(logIn[:n]), "|")
	name := message[0]
	//password := message[1]
	return name, true
}

//处理请求函数
func process(conn net.Conn, DB *sql.DB) {
	defer conn.Close() //关闭连接
	name, ok := LegalAccount(conn)
	if ok == false {
		_, err := conn.Write([]byte("No"))
		if err != nil {
			fmt.Println("返回验证信息错误,err = ", err)
			return
		}
	}
	fmt.Println(name, "已经登陆")
	_, err := conn.Write([]byte("OK"))
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf) //接受数据
		if err != nil {
			fmt.Println(name, "退出登录")
			break
		}
		message := string(buf[:n]) //接受的数据
		fmt.Println(name, ":", message)
		_, err = conn.Write([]byte("OK"))
		if err != nil {
			fmt.Println("返回验证信息错误,err = ", err)
			return
		}
	}
}

// 主服务器
func main() {
	//监听
	DB, err := DBInit()
	if err != nil {
		fmt.Println("数据库连接失败,err = ", err)
		return
	}
	lister, err := net.Listen("tcp", "127.0.0.1:8000") //监听此端口
	if err != nil {
		fmt.Println(err)
		return
	}
	defer lister.Close() //监听关闭
	//阻塞等待用户
	for {
		conn, err := lister.Accept() //返回连接
		if err != nil {
			fmt.Println(err)
			continue
		}
		//处理请求
		go process(conn, DB)
	}
}

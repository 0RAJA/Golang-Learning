package main

import (
	"fmt"
	"log"
	"os"
)

//Go语言内置的log包实现了简单的日志服务
/*
1.使用logger
log包定义了Logger类型，该类型提供了一些格式化输出的方法。
本包也提供了一个预定义的“标准”logger，可以通过调用函数
	Print系列(Print|Printf|Println）、
	Fatal系列（Fatal|Fatalf|Fatalln）、
	Panic系列（Panic|Panicf|Panicln）来使用。
logger会打印每条日志信息的日期、时间，默认输出到系统的标准错误。Fatal系列函数会在写入日志信息后调用os.Exit(1)。Panic系列函数会在写入日志信息后panic。
2.配置logger
    func Flags() int
    func SetFlags(flag int)
log标准库中的Flags函数会返回标准logger的输出配置，而SetFlags函数用来设置标准logger的输出配置。
const (
    // 控制输出日志信息的细节，不能控制输出的顺序和格式。
    // 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
    Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
    Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
    LUTC                          // 使用UTC时间
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
)
3.配置日志前缀
    func Prefix() string
    func SetPrefix(prefix string)
其中Prefix函数用来查看标准logger的输出前缀，SetPrefix函数用来设置输出前缀。
4.配置日志输出位置
	func SetOutput(w io.Writer)
SetOutput函数用来设置标准logger的输出目的地，默认是标准错误输出。
5.创建logger
	func New(out io.Writer, prefix string, flag int) *Logger
*/
func main() {
	if false {
		log.Println("One")
		v := "Two"
		log.Printf("%s", v)
		log.Fatalln("fatal")
		log.Panicln("panic")
	}
	if true {
		log.Println("en...")
		fmt.Println(log.Flags()) //Flags 返回标准记录器的输出标志。 标志位是Ldate 、 Ltime等。
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		log.Println("ha...")
	}
	if false {
		log.Println("old")
		log.SetPrefix("[System]")
		log.Println("New")
	}
	if false {
		logFile, err := os.OpenFile("src/package/log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0744)
		if err != nil {
			log.Println(err)
			return
		}
		defer logFile.Close()
		log.SetOutput(logFile)
		log.SetPrefix("[Raja]")
		log.Println("这是一条很普通的日志。")
	}
	if false {
		logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
		logger.Println("这是自定义的logger记录的日志。")
	}
}

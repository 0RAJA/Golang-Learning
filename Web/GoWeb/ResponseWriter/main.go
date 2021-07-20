package main

import (
	"encoding/json"
	"net/http"
)

/*
type ResponseWriter interface {
    // Header返回一个Header类型值，该值会被WriteHeader方法发送。
    // 在调用WriteHeader或Write方法后再改变该对象是没有意义的。
    Header() Header
    // WriteHeader该方法发送HTTP回复的头域和状态码。
    // 如果没有被显式调用，第一次调用Write时会触发隐式调用WriteHeader(http.StatusOK)
    // WriterHeader的显式调用主要用于发送错误码。
    WriteHeader(int)
    // Write向连接中写入作为HTTP的一部分回复的数据。
    // 如果被调用时还未调用WriteHeader，本方法会先调用WriteHeader(http.StatusOK)
    // 如果Header中没有"Content-Type"键，
    // 本方法会使用包函数DetectContentType检查数据的前512字节，将返回值作为该键的值。
    Write([]byte) (int, error)
}
*/
func handleConn1(w http.ResponseWriter, r *http.Request) {
	//响应字符串
	w.Write([]byte("OK"))
}
func handleConn2(w http.ResponseWriter, r *http.Request) {
	//响应html页面

	html := `<html>
<head>
	<title>测试响应内容为网页</title>
	<meta charset="utf-8"/>
</head>
<body>
	HELLO!!!
</body>
</html>`
	w.Write([]byte(html))
}

func handleConn3(w http.ResponseWriter, r *http.Request) {
	//设置响应头中内容的类型
	w.Header().Set("Content-Type", "application/json")
	stu := struct {
		Name string
		Age  int
	}{Name: "ZS", Age: 19}
	js, _ := json.Marshal(stu)
	w.Write(js)
}
func handleConn4(w http.ResponseWriter, r *http.Request) {
	//重定向
	w.Header().Set("Location", "https:www.baidu.com")
	w.WriteHeader(302)
}
func main() {
	http.HandleFunc("/go", handleConn4)
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"net/http"
	"text/template"
)

/*
type Dir string
	Dir使用限制到指定目录树的本地文件系统实现了http.FileSystem接口。空Dir被视为"."，即代表当前目录。
func FileServer(root FileSystem) Handler
	FileServer返回一个使用FileSystem接口root提供文件访问服务的HTTP处理器。要使用操作系统的FileSystem接口实现，可使用http.Dir：
func StripPrefix(prefix string, h Handler) Handler
	StripPrefix返回一个处理器，该处理器会将请求的URL.Path字段中给定前缀prefix去除后再交由h处理。StripPrefix会向URL.Path字段中没有给定前缀的请求回复404 page not found。
举例:
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
*/
//IndexHandle 去首页

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t := template.Must(template.ParseFiles("src/Web/GoWeb/bookstore/views/index.html"))
	t.Execute(w, "")
}

func main() {
	// 处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("src/Web/GoWeb/bookstore/views/static"))))
	//将访问带/static/前缀的地址 更改为在 src/Web/GoWeb/bookstore/views/static下寻找相关文件
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("src/Web/GoWeb/bookstore/views/pages"))))
	//处理主页
	http.HandleFunc("/main", IndexHandle)

	http.ListenAndServe(":8080", nil)
}

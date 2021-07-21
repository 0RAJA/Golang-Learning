package main

import (
	"Web/GoWeb/bookstore/controller"
	"net/http"
	"text/template"
)

//IndexHandle 去首页
func IndexHandle(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t := template.Must(template.ParseFiles("src/Web/GoWeb/bookstore/views/index.html"))
	t.Execute(w, "")
}

func main() {
	// 处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("src/Web/GoWeb/bookstore/views/static"))))
	//直接去HTML页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("src/Web/GoWeb/bookstore/views/pages"))))
	http.HandleFunc("/main", IndexHandle)
	//去登录
	http.HandleFunc("/login", controller.Login)
	//去注册
	http.HandleFunc("/regist", controller.Regist)
	//通过ajax验证用户名是否可用
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	//获取所有图书
	http.HandleFunc("/getBooks", controller.GetBooks)
	//添加图书
	http.HandleFunc("/addAndModify", controller.AddAndModify)
	//删除图书
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	//修改图书
	http.HandleFunc("/addOrModify", controller.AddOrModify)
	//监听
	http.ListenAndServe(":8080", nil)
}

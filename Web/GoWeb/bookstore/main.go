package main

import (
	"Web/GoWeb/bookstore/controller"
	"net/http"
)

func main() {
	// 处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("src/Web/GoWeb/bookstore/views/static"))))
	//直接去HTML页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("src/Web/GoWeb/bookstore/views/pages"))))
	//去首页
	http.HandleFunc("/main", controller.GetPageBooksByPrice)
	//去登录
	http.HandleFunc("/login", controller.Login)
	//去注销
	http.HandleFunc("/logout", controller.Logout)
	//去注册
	http.HandleFunc("/regist", controller.Regist)
	//通过ajax验证用户名是否可用
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	//获取所有图书
	//http.HandleFunc("/getBooks", controller.GetBooks)
	//获取分页图书
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)
	//添加图书
	http.HandleFunc("/addAndModify", controller.AddAndModify)
	//删除图书
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	//修改图书
	http.HandleFunc("/addOrModify", controller.AddOrModify)
	//添加图书到购物车
	http.HandleFunc("/addBookToCart", controller.AddBookToCart)
	//获取购物车信息
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)
	//清空购物车
	http.HandleFunc("/deleteCart", controller.DeleteCart)
	//删除购物项
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)
	//更新购物项
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)
	//监听
	http.ListenAndServe(":8080", nil)
}

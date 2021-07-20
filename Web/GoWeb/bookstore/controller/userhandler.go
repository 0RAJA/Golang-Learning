package controller

import (
	"Web/GoWeb/bookstore/dao"
	"Web/GoWeb/bookstore/model"
	"net/http"
	"text/template"
)

//Login 处理用户登录函数
func Login(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username, password := r.PostFormValue("username"), r.FormValue("password")
	//验证用户名和密码
	_, err := dao.CheckUserNamePwd(username, password)
	if err != nil { //登录失败
		t := template.Must(template.ParseFiles("src/Web/GoWeb/bookstore/views/pages/user/login.html"))
		t.Execute(w, "用户名或密码错误")
	} else {
		t := template.Must(template.ParseFiles("src/Web/GoWeb/bookstore/views/pages/user/login_success.html"))
		t.Execute(w, username)
	}
}

func Regist(w http.ResponseWriter, r *http.Request) {
	//获取表单信息
	username, password, email := r.PostFormValue("username"), r.FormValue("password"), r.FormValue("email")
	//注册信息
	user := &model.User{
		UserName: username,
		Password: password,
		Email:    email,
	}
	err := dao.SaveUser(user)
	if err != nil {
		t := template.Must(template.ParseFiles("src/Web/GoWeb/bookstore/views/pages/user/regist.html"))
		t.Execute(w, "用户名已存在")
	} else {
		t := template.Must(template.ParseFiles("src/Web/GoWeb/bookstore/views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}

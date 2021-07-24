package controller

import (
	"Web/GoWeb/bookstore/dao"
	"Web/GoWeb/bookstore/model"
	"fmt"
	"github.com/gofrs/uuid"
	"net/http"
	"text/template"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	//获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取cookie的value值
		cookieValue := cookie.Value
		//删除数据库中的session
		dao.DeleteSession(cookieValue)
		//设置cookie失效
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
	//去首页
	GetPageBooksByPrice(w, r)
}

//Login 处理用户登录函数
func Login(w http.ResponseWriter, r *http.Request) {
	flag, _ := dao.IsLogin(r)
	if flag == true {
		//登陆过了
		GetPageBooksByPrice(w, r)
		return
	}
	//获取用户名和密码
	username, password := r.PostFormValue("username"), r.FormValue("password")
	//验证用户名和密码
	user, err := dao.CheckUserNamePwd(username, password)
	if err != nil { //登录失败
		t := template.Must(template.ParseFiles("src/Web/GoWeb/bookstore/views/pages/user/login.html"))
		t.Execute(w, "用户名或密码错误")
	} else {
		//生成uuid作为session的id
		sessionID, _ := uuid.NewV4()
		session := model.Session{
			SessionID: sessionID.String(),
			UserName:  user.UserName,
			UserID:    user.ID,
		}
		//将uuid保存到数据库中
		err := dao.AddSession(&session)
		if err != nil {
			fmt.Println(err)
		}
		//创建cookie与session关联
		cookie := http.Cookie{
			Name:     "user",
			Value:    sessionID.String(),
			HttpOnly: true,
		}
		//将cookie发给浏览器
		http.SetCookie(w, &cookie)
		t := template.Must(template.ParseFiles("src/Web/GoWeb/bookstore/views/pages/user/login_success.html"))
		t.Execute(w, user)
	}
}

//Regist 注册
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
		t.Execute(w, user)
	}
}

// CheckUserName 验证用户名
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	//获取用户名
	username := r.FormValue("username")
	_, err := dao.CheckUserName(username)
	if err != nil {
		w.Write([]byte("OK"))
	} else {
		w.Write([]byte("NO"))
	}
}

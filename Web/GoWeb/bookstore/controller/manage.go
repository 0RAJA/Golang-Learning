package controller

import (
	"net/http"
	"text/template"
)

// Manage 管理员界面
func Manage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("src/Web/GoWeb/bookstore/views/pages/manager/manager.html"))
	_ = t.Execute(w, "")
}

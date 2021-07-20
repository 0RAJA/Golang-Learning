package main

import (
	"net/http"
	"text/template"
)

type Person struct {
	Name  string
	Email string
}

func handleConn(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("src/Web/GoWeb/动作/main.html"))
	//slice := []*Person{{Name: "张三", Email: "123@qq.com"}, {Name: "李四", Email: "234@qq.com"}, {Name: "王麻子", Email: "567@qq.com"}}
	mMap := map[string]string{"a": "aa", "b": "bb", "c": "cc"}
	t.Execute(w, mMap)
}

func main() {
	http.HandleFunc("/handle", handleConn)
	http.ListenAndServe(":8080", nil)
}

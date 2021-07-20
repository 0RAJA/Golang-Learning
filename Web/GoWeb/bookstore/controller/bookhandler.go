package controller

import (
	"Web/GoWeb/bookstore/dao"
	"Web/GoWeb/bookstore/model"
	"net/http"
	"strconv"
	"text/template"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	//获取图书
	books, _ := dao.GetBooks()
	//解析模板文件
	t := template.Must(template.ParseFiles("src/Web/GoWeb/bookstore/views/pages/manager/book_manager.html"))
	//执行
	t.Execute(w, books)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	book := model.Book{
		Title:   r.FormValue("title"),
		Author:  r.FormValue("author"),
		ImgPath: "src/Web/GoWeb/bookstore/views/static/img/default.jpg",
	}
	book.Price, _ = strconv.ParseFloat(r.FormValue("price"), 10)
	book.Sales, _ = strconv.Atoi(r.FormValue("sales"))
	book.Stock, _ = strconv.Atoi(r.FormValue("stock"))
	_ = dao.AddBook(&book)
	GetBooks(w, r)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID, _ := strconv.Atoi(r.FormValue("bookID"))
	_ = dao.DeleteBook(bookID)
	GetBooks(w, r)
}

func GetModifyID(w http.ResponseWriter, r *http.Request) {
	bookID, _ := strconv.Atoi(r.FormValue("bookID"))
	book, _ := dao.GetBookByID(bookID)
	t := template.Must(template.ParseFiles("src/Web/GoWeb/bookstore/views/pages/manager/book_modify.html"))
	_ = t.Execute(w, book)
}

func Modify(w http.ResponseWriter, r *http.Request) {
	book := model.Book{
		Title:  r.FormValue("title"),
		Author: r.FormValue("author"),
	}
	book.ID, _ = strconv.Atoi(r.FormValue("bookID"))
	book.Price, _ = strconv.ParseFloat(r.FormValue("price"), 10)
	book.Sales, _ = strconv.Atoi(r.FormValue("sales"))
	book.Stock, _ = strconv.Atoi(r.FormValue("stock"))
	_ = dao.Modify(&book)
	GetBooks(w, r)
}

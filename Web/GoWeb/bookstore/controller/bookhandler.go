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

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID, _ := strconv.Atoi(r.FormValue("bookID"))
	_ = dao.DeleteBook(bookID)
	GetBooks(w, r)
}

func AddOrModify(w http.ResponseWriter, r *http.Request) {
	bookID, _ := strconv.Atoi(r.FormValue("bookID"))
	book, _ := dao.GetBookByID(bookID)
	t := template.Must(template.ParseFiles("src/Web/GoWeb/bookstore/views/pages/manager/book_edit.html"))
	if bookID > 0 {
		_ = t.Execute(w, book)
	} else {
		_ = t.Execute(w, "")
	}
}

func AddAndModify(w http.ResponseWriter, r *http.Request) {
	book := model.Book{
		Title:  r.FormValue("title"),
		Author: r.FormValue("author"),
	}
	book.Price, _ = strconv.ParseFloat(r.FormValue("price"), 10)
	book.Sales, _ = strconv.Atoi(r.FormValue("sales"))
	book.Stock, _ = strconv.Atoi(r.FormValue("stock"))
	book.ID, _ = strconv.Atoi(r.FormValue("bookID"))
	if book.ID > 0 {
		_ = dao.Modify(&book)
	} else {
		book.ImgPath = "src/Web/GoWeb/bookstore/views/static/img/default.jpg"
		_ = dao.AddBook(&book)
	}
	GetBooks(w, r)
}


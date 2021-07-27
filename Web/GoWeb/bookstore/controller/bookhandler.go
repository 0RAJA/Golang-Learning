package controller

import (
	"Web/GoWeb/bookstore/dao"
	"Web/GoWeb/bookstore/model"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

const DirPATH = "src/Web/GoWeb/bookstore/views/static/img"

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// GetPageBooks 获取分页图书
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	if r.FormValue("pageNo") == "" {
		pageNo = 1
	}
	//获取分页图书
	page, _ := dao.GetPageBooks(pageNo)
	//解析模板文件
	t := template.Must(template.ParseFiles("src/Web/GoWeb/bookstore/views/pages/manager/book_manager.html"))
	//执行
	_ = t.Execute(w, page)
}

// GetPageBooksByPrice 通过价格筛选图书
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo, _ := strconv.Atoi(r.FormValue("pageNo"))
	if r.FormValue("pageNo") == "" {
		pageNo = 1
	}
	var page *model.Page
	min, max := r.FormValue("min"), r.FormValue("max")
	if min == "" || max == "" {
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		min, _ := strconv.ParseFloat(min, 10)
		max, _ := strconv.ParseFloat(max, 10)
		page, _ = dao.GetPageBooksByPrice(pageNo, min, max)
	}
	page.Min, page.Max = min, max
	ok, session := dao.IsLogin(r)
	//执行
	if ok {
		//已经登陆
		page.IsLogin = true
		page.UserName = session.UserName
	}
	t := template.Must(template.ParseFiles("src/Web/GoWeb/bookstore/views/index.html"))
	_ = t.Execute(w, page)
}

// DeleteBook 删除图书
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID, _ := strconv.Atoi(r.FormValue("bookID"))
	err := dao.DeleteBook(bookID)
	if err != nil {
		log.Println(err)
		return
	}
	GetPageBooks(w, r)
}

// AddOrModify 修改或者增加
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

// AddAndModify 增加且修改
func AddAndModify(w http.ResponseWriter, r *http.Request) {
	book := model.Book{
		Title:  r.FormValue("title"),
		Author: r.FormValue("author"),
	}
	book.Price, _ = strconv.ParseFloat(r.FormValue("price"), 10)
	book.Sales, _ = strconv.Atoi(r.FormValue("sales"))
	book.Stock, _ = strconv.Atoi(r.FormValue("stock"))
	book.ID, _ = strconv.Atoi(r.FormValue("bookID"))
	f, h, err := r.FormFile("imgPath")
	if err != nil {
		log.Println(f, h, err)
	}
	defer f.Close()
	fileName := h.Filename
	t, err := os.Create(DirPATH + "/" + fileName)
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	defer t.Close()
	if _, err := io.Copy(t, f); err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	book.ImgPath = "static/img/" + fileName
	if book.ID > 0 {
		_ = dao.Modify(&book)
	} else {
		_ = dao.AddBook(&book)
	}
	GetPageBooks(w, r)
}

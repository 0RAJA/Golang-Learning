package dao

import (
	"Web/GoWeb/bookstore/model"
	"fmt"
	"testing"
)

func TestBooks(t *testing.T) {
	t.Run("测试获取图书:", TestGetBooks)
}

func TestGetBooks(t *testing.T) {
	books, err := GetBooks()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, v := range books {
			fmt.Println(*v)
		}
	}
}
func TestAddBook(t *testing.T) {
	book := model.Book{
		Title:   "1",
		Author:  "1",
		Price:   1,
		Sales:   1,
		Stock:   1,
		ImgPath: "111",
	}
	err := AddBook(&book)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("OK")
}

func TestDeleteBook(t *testing.T) {
	_ = DeleteBook(33)
}

func TestGetBookByID(t *testing.T) {
	book, err := GetBookByID(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*book)
}

func TestModify(t *testing.T) {
	book := model.Book{
		ID:      35,
		Title:   "2",
		Author:  "2",
		Price:   2,
		Sales:   2,
		Stock:   2,
		ImgPath: "2",
	}
	err := Modify(&book)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("OK")
}

func TestGetPageBooks(t *testing.T) {
	page, err := GetPageBooks(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < len(page.Books); i++ {
		fmt.Println(*page.Books[i])
	}
	fmt.Println(page)
}

func TestGetPageBooksByPrice(t *testing.T) {
	page, err := GetPageBooksByPrice(1, 19, 20)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < len(page.Books); i++ {
		fmt.Println(*page.Books[i])
	}
	fmt.Println(page)
}

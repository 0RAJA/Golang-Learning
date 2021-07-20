package dao

import (
	"Web/GoWeb/bookstore/model"
	"Web/GoWeb/bookstore/utils"
)

//GetBooks 获取数据库中所有的图书
func GetBooks() ([]*model.Book, error) {
	sqlStr := "SELECT * FROM books"
	rows, err := utils.DB.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, nil
}

//AddBook 向数据库中添加一本图书
func AddBook(book *model.Book) error {
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	_, err := utils.DB.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

// DeleteBook 从数据库中删除图书
func DeleteBook(bookID int) error {
	sqlStr := "delete from books where `id` = ?"
	_, err := utils.DB.Exec(sqlStr, bookID)
	if err != nil {
		return err
	}
	return nil
}

func GetBookByID(bookID int) (*model.Book, error) {
	sqlStr := "select * from books where `id` = ?"
	var book model.Book
	row := utils.DB.QueryRow(sqlStr, bookID)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func Modify(book *model.Book) error {
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=? where `id`=?"
	_, err := utils.DB.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ID)
	if err != nil {
		return err
	}
	return nil
}

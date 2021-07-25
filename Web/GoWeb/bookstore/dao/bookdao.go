package dao

import (
	"Web/GoWeb/bookstore/model"
	"Web/GoWeb/bookstore/utils"
)

//GetBooks 获取所有的图书
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

//AddBook 添加一本图书
func AddBook(book *model.Book) error {
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	_, err := utils.DB.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

// DeleteBook 删除图书
func DeleteBook(bookID int) error {
	sqlStr := "delete from books where `id` = ?"
	_, err := utils.DB.Exec(sqlStr, bookID)
	if err != nil {
		return err
	}
	return nil
}

// GetBookByID 通过图书ID获取图书信息
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

// Modify 更新图书信息
func Modify(book *model.Book) error {
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=? where `id`=?"
	_, err := utils.DB.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetPageBooks 通过pageNo获取当前页数的图书信息
func GetPageBooks(pageNo int) (*model.Page, error) {
	page := model.Page{
		PageNo:   pageNo,
		PageSize: 4,
	}
	//获取总记录数TotalPageNo和TotalRecord
	sqlStr1 := "select count(*) from books"
	err := utils.DB.QueryRow(sqlStr1).Scan(&page.TotalRecord)
	if err != nil {
		return nil, err
	}
	page.TotalPageNo = page.TotalRecord / page.PageSize
	if page.TotalRecord%page.PageSize != 0 {
		page.TotalPageNo++
	}
	//通过limit获取图书
	//公式: F(pageNo) = (pageNo-1)*pageSize
	sqlStr2 := "select * from books limit ?,?"
	rows, err := utils.DB.Query(sqlStr2, (page.PageNo-1)*page.PageSize, page.PageSize)
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
	page.Books = books
	return &page, nil
}

// GetPageBooksByPrice 通过pageNo获取当前页数的图书信息
func GetPageBooksByPrice(pageNo int, min, max float64) (*model.Page, error) {
	page := model.Page{
		PageNo:   pageNo,
		PageSize: 4,
	}
	//获取总记录数TotalPageNo和TotalRecord
	sqlStr1 := "select count(*) from books where price between ? and ?"
	err := utils.DB.QueryRow(sqlStr1, min, max).Scan(&page.TotalRecord)
	if err != nil {
		return nil, err
	}
	page.TotalPageNo = page.TotalRecord / page.PageSize
	if page.TotalRecord%page.PageSize != 0 {
		page.TotalPageNo++
	}
	//通过limit获取图书
	//公式: F(pageNo) = (pageNo-1)*pageSize
	sqlStr2 := "select * from books where price between ? and ? limit ?,?"
	rows, err := utils.DB.Query(sqlStr2, min, max, (page.PageNo-1)*page.PageSize, page.PageSize)
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
	page.Books = books
	return &page, nil
}

// UpdateBook 更新图书信息
func UpdateBook(book *model.Book) error {
	sql := "update books set title = ?,author=?,price=?,sales=?,stock = ? where id = ?"
	_, err := utils.DB.Exec(sql, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ID)
	if err != nil {
		return err
	}
	return nil
}

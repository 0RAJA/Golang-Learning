package dao

import (
	"Web/GoWeb/bookstore/model"
	"Web/GoWeb/bookstore/utils"
)

// AddCartItem 添加到购物车
func AddCartItem(item *model.CartItem) error {
	sqlStr := "insert into cart_items(`count`,amount,book_id,cart_id)  values(?,?,?,?)"
	_, err := utils.DB.Exec(sqlStr, item.Count, item.GetAmount(), item.Book.ID, item.CartID)
	if err != nil {
		return err
	}
	return nil
}

// GetCartItemByBookIDAndCardID 根据图书ID和购物车ID获取对应的购物项
func GetCartItemByBookIDAndCardID(bookID int, cardID string) (*model.CartItem, error) {
	sqlStr := "select id,`count`,amount,cart_id from cart_items where book_id = ? and cart_id = ?"
	var cartItem model.CartItem
	err := utils.DB.QueryRow(sqlStr, bookID, cardID).Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil, err
	}
	cartItem.Book, _ = GetBookByID(bookID)
	return &cartItem, nil
}

// UpdateBookCountAndAmount 更新购物车中书的数量以及价格
func UpdateBookCountAndAmount(item *model.CartItem) error {
	sqlStr := "update cart_items set count = ?,amount = ? where book_id = ? and cart_id = ?"
	_, err := utils.DB.Exec(sqlStr, item.Count, item.GetAmount(), item.Book.ID, item.CartID)
	if err != nil {
		return err
	}
	return nil
}

// GetCartItemsByCartID 根据购物车ID获取购物车中所有购物项
func GetCartItemsByCartID(cartID string) ([]*model.CartItem, error) {
	sqlStr := "select id,`count`,amount,cart_id,book_id from cart_items where cart_id = ?"
	rows, err := utils.DB.Query(sqlStr, cartID)
	if err != nil {
		return nil, err
	}
	var cartItems []*model.CartItem
	for rows.Next() {
		var cartItem model.CartItem
		var bookID int
		err := rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID, &bookID)
		if err != nil {
			return nil, err
		}
		cartItem.Book, _ = GetBookByID(bookID)
		cartItems = append(cartItems, &cartItem)
	}
	return cartItems, nil
}

// DeleteCartItemByCartID 根据购物车ID删除所有的购物项
func DeleteCartItemByCartID(cartID string) error {
	sqlStr := "delete from cart_items where cart_id = ?"
	_, err := utils.DB.Exec(sqlStr, cartID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCartItemByID 通过ItemID删除购物项
func DeleteCartItemByID(cartItemID int) error {
	sqlStr := "delete from cart_items where id = ?"
	_, err := utils.DB.Exec(sqlStr, cartItemID)
	if err != nil {
		return err
	}
	return nil
}

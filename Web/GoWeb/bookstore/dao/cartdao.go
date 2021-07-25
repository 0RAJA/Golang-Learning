package dao

import (
	"Web/GoWeb/bookstore/model"
	"Web/GoWeb/bookstore/utils"
)

// AddCart 添加购物车到数据库
func AddCart(cart *model.Cart) error {
	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"
	_, err := utils.DB.Exec(sqlStr, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		return err
	}
	//遍历购物车中的购物项
	for _, cartItem := range cart.CartItems {
		//将购物项插入到数据库中
		err := AddCartItem(cartItem)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetCartByUserID 根据用户ID查询对应的购物车
func GetCartByUserID(userID int) (*model.Cart, error) {
	sqlStr := "select * from carts where user_id = ?"
	var cart model.Cart
	err := utils.DB.QueryRow(sqlStr, userID).Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	if err != nil {
		return nil, err
	}
	cartItems, err := GetCartItemsByCartID(cart.CartID)
	if err != nil {
		return nil, err
	}
	cart.CartItems = cartItems
	return &cart, nil
}

// UpdateCart 更新购物车中的图书的总数量以及总金额
func UpdateCart(cart *model.Cart) error {
	sqlStr := "update carts set total_count = ?,total_amount = ? where id = ?"
	_, err := utils.DB.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCartByCartID 根据购物车ID删除购物车
func DeleteCartByCartID(cartID string) error {
	//删除购物车之前需要先删除购物项
	err := DeleteCartItemByCartID(cartID)
	if err != nil {
		return err
	}
	sqlStr := "delete from carts where id = ?"
	_, err = utils.DB.Exec(sqlStr, cartID)
	if err != nil {
		return err
	}
	return nil
}

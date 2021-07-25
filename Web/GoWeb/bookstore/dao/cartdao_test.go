package dao

import (
	"Web/GoWeb/bookstore/model"
	"fmt"
	"testing"
)

func TestAddCart(t *testing.T) {
	//设置要买的第一本书
	book1, _ := GetBookByID(2)
	book2, _ := GetBookByID(3)
	//购物项
	cartItem1 := &model.CartItem{Book: book1, Count: 2, CartID: "668"}
	cartItem2 := &model.CartItem{Book: book2, Count: 2, CartID: "668"}
	cartItems := []*model.CartItem{cartItem1, cartItem2}
	cart := model.Cart{
		CartID:    "668",
		CartItems: cartItems,
		UserID:    2,
	}
	err := AddCart(&cart)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestGetCartByUserID(t *testing.T) {
	cart, err := GetCartByUserID(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cart)
	for i := 0; i < len(cart.CartItems); i++ {
		fmt.Println(*cart.CartItems[i])
	}
}

func TestDeleteCartByCartID(t *testing.T) {
	err := DeleteCartByCartID("39ee5efe-c71d-482b-92ce-b3d6469f512a")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestDeleteCartItemByID(t *testing.T) {
	err := DeleteCartItemByID(68)
	if err != nil {
		return
	}
}

package controller

import (
	"Web/GoWeb/bookstore/dao"
	"Web/GoWeb/bookstore/model"
	"fmt"
	"github.com/gofrs/uuid"
	"net/http"
	"strconv"
)

// AddBookToCart 添加图书到购物车
func AddBookToCart(w http.ResponseWriter, r *http.Request) {
	//获取要添加图书的id
	bookIDStr := r.FormValue("bookID")
	bookID, _ := strconv.Atoi(bookIDStr)
	//根据图书ID获取图书信息
	book, _ := dao.GetBookByID(bookID)
	//判断是否登录
	flag, session := dao.IsLogin(r)
	if flag == true {
		//获取用户ID
		userID := session.UserID
		//获取购物车
		cart, err := dao.GetCartByUserID(userID)
		if err != nil { //没有购物车
			cartID, _ := uuid.NewV4()
			//创建购物车
			cart = &model.Cart{
				CartID: cartID.String(),
				UserID: userID,
			}
			//创建购物项
			var cartItems []*model.CartItem
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartID: cartID.String(),
			}
			//将购物项添加到切片中
			cartItems = append(cartItems, cartItem)
			//将购物项切片设置到购物车中
			cart.CartItems = cartItems
			err = dao.AddCart(cart)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			//当前用户已经有购物车.判断有没有此书
			cartItem, err := dao.GetCartItemByBookIDAndCardID(bookID, cart.CartID)
			if err != nil { //找不到该图书说明要创建
				//创建购物项
				cartItem = &model.CartItem{
					Book:   book,
					Count:  1,
					CartID: cart.CartID,
				}
				//将购物项添加到cart切片中
				cart.CartItems = append(cart.CartItems, cartItem)
				//将购物项切片
				err := dao.AddCartItem(cartItem)
				if err != nil {
					fmt.Println(err)
					return
				}
			} else { //找到该图书,只需要把此书的数量+1即可
				for _, v := range cart.CartItems {
					if v.Book.ID == cartItem.Book.ID {
						v.Count++
						err := dao.UpdateBookCount(v.Count, v.Book.ID, v.CartID)
						if err != nil {
							fmt.Println(err)
							return
						}
					}
				}
			}
			//不管之前是否有没有购买此书,最后都要更新购物车中图书数量和金额
			err = dao.UpdateCart(cart)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		_, _ = w.Write([]byte(book.Title))
	} else {

	}
}

package controller

import (
	"Web/GoWeb/bookstore/dao"
	"Web/GoWeb/bookstore/model"
	"encoding/xml"
	"github.com/gofrs/uuid"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

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
				log.Println(err)
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
					log.Println(err)
					return
				}
			} else { //找到该图书,只需要把此书的数量+1即可
				for _, v := range cart.CartItems {
					if v.Book.ID == cartItem.Book.ID {
						v.Count++
						v.Amount = v.GetAmount()
						err := dao.UpdateBookCountAndAmount(v)
						if err != nil {
							log.Println(err)
							return
						}
					}
				}
			}
			//不管之前是否有没有购买此书,最后都要更新购物车中图书数量和金额
			err = dao.UpdateCart(cart)
			if err != nil {
				log.Println(err)
				return
			}
		}
		_, _ = w.Write([]byte(book.Title))
	} else {

	}
}

func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	flag, session := dao.IsLogin(r)
	if flag == true { //登录了
		cart, err := dao.GetCartByUserID(session.UserID)
		if err != nil {
			cart = &model.Cart{HaveShop: false}
		} else { //有购物车
			cart.HaveShop = true
		}
		//给个名字
		cart.UserName = session.UserName
		//解析模板
		t := template.Must(template.ParseFiles("src/Web/GoWeb/bookstore/views/pages/cart/cart.html"))
		// 执行
		_ = t.Execute(w, cart)
	} else { //去登录
		Login(w, r)
	}
}

// DeleteCart 清空购物车
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	cartID := r.FormValue("cartID")
	err := dao.DeleteCartByCartID(cartID)
	if err != nil {
		log.Println(err)
		return
	}
	GetCartInfo(w, r)
}

// DeleteCartItem 删除购物项
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	//获取要删除的购物项的ID
	cartItemID, _ := strconv.Atoi(r.FormValue("cartItemID"))
	//获取session
	flag, session := dao.IsLogin(r)
	if flag {
		//获取用户ID
		userID := session.UserID
		//获取用户购物车
		cart, err := dao.GetCartByUserID(userID)
		if err != nil {
			log.Println(err)
			return
		}
		//获取购物车中的购物项
		var cartItems = cart.CartItems
		for k, v := range cartItems {
			if v.CartItemID == cartItemID {
				//将此切片从购物车中移除
				cartItems = append(cartItems[:k], cartItems[k+1:]...)
				err := dao.DeleteCartItemByID(cartItemID)
				if err != nil {
					log.Println(err)
					return
				}
				cart.CartItems = cartItems //因为地址变了,得重新赋值下
				break
			}
		}
		//更新购物车中的总数量和金额
		err = dao.UpdateCart(cart)
		if err != nil {
			log.Println(err)
			return
		}
		//重新查询购物车信息
		GetCartInfo(w, r)
	}
}

type update struct {
	ItemID      int
	Amount      float64
	TotalCount  int
	TotalAmount float64
}

// UpdateCartItem 更新购物项
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	//获取要更新的购物项的ID
	cartItemID, _ := strconv.Atoi(r.FormValue("cartItemID"))
	bookCount, _ := strconv.Atoi(r.FormValue("bookCount"))
	//获取session
	flag, session := dao.IsLogin(r)
	if flag {
		//获取用户ID
		userID := session.UserID
		//获取用户购物车
		cart, err := dao.GetCartByUserID(userID)
		if err != nil {
			log.Println(err)
			return
		}
		updateMessage := update{
			ItemID: cartItemID,
		}
		//获取购物车中的购物项
		var cartItems = cart.CartItems
		for _, v := range cartItems {
			if v.CartItemID == cartItemID {
				v.Count = bookCount
				v.Amount = float64(v.Count) * v.Book.Price
				//更新数据库中改购物项的图书的数量和金额
				err = dao.UpdateBookCountAndAmount(v)
				if err != nil {
					log.Println(err)
					return
				}
				updateMessage.Amount = v.Amount
				break
			}
		}
		//更新数据库中的总数量和金额
		err = dao.UpdateCart(cart)
		if err != nil {
			log.Println(err)
			return
		}
		updateMessage.TotalCount = cart.GetTotalCount()
		updateMessage.TotalAmount = cart.GetTotalAmount()
		//回传信息
		w.Header().Set("Content-Type", "text/xml;charset=utf-8")
		v, _ := xml.Marshal(&updateMessage)
		_, _ = w.Write(v)
	}
}

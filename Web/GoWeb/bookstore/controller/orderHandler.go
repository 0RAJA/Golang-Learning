package controller

import (
	"Web/GoWeb/bookstore/dao"
	"Web/GoWeb/bookstore/model"
	"github.com/gofrs/uuid"
	"log"
	"net/http"
	"text/template"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// Checkout 去结账
func Checkout(w http.ResponseWriter, r *http.Request) {
	//获取购物车内容session
	flag, session := dao.IsLogin(r)
	if flag == true { // 登录成功
		//获取userID
		userID := session.UserID
		//获取cart
		cart, err := dao.GetCartByUserID(userID)
		if err != nil {
			log.Println(err)
			return
		}
		//生成订单号
		orderID, _ := uuid.NewV4()
		//生成订单
		order := model.Order{
			OrderID:     orderID.String(),
			CreateTime:  time.Now().Format("2006-1-2 15:04:05 MST Mon Jan"),
			TotalCount:  cart.TotalCount,
			TotalAmount: cart.TotalAmount,
			State:       0,
			UserID:      userID,
		}
		//保存订单
		err = dao.AddOrder(&order)
		if err != nil {
			log.Println(err)
			return
		}
		//保存订单项
		for _, v := range cart.CartItems {
			//创建订单项
			orderItem := model.OrderItem{
				Count:   v.Count,
				Amount:  v.Amount,
				Title:   v.Book.Title,
				Author:  v.Book.Author,
				Price:   v.Book.Price,
				ImgPath: v.Book.ImgPath,
				OrderID: orderID.String(),
			}
			//保存订单项
			err := dao.AddOrderItem(&orderItem)
			if err != nil {
				log.Println(err)
				return
			}
			//更新图书库存和销量
			book := v.Book
			book.Sales += orderItem.Count
			book.Stock -= orderItem.Count
			err = dao.UpdateBook(book)
			if err != nil {
				log.Println(err)
				return
			}
		}
		//清空购物车
		err = dao.DeleteCartByCartID(cart.CartID)
		if err != nil {
			log.Println(err)
			return
		}
		//添加订单号
		session.OrderID = orderID.String()
		//刷新页面
		t := template.Must(template.ParseFiles("src/Web/GoWeb/bookstore/views/pages/cart/checkout.html"))
		_ = t.Execute(w, session)
	} else {
		Login(w, r)
	}
}

// GetAllOrders 获取所有订单
func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	//获取订单
	orders, err := dao.GetAllOrders()
	if err != nil {
		log.Println(err)
		return
	}
	for _, v := range orders {
		v.NoSend, v.NoSend, v.SendComplete = false, false, false
		switch v.State {
		case 0:
			v.NoSend = true
		case 1:
			v.SendComplete = true
		case 2:
			v.Complete = true
		}
	}
	//解析模板
	t := template.Must(template.ParseFiles("src/Web/GoWeb/bookstore/views/pages/order/order_manager.html"))
	_ = t.Execute(w, orders)
}

// GetOrderInfo 获取订单详细信息
func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	//获取订单号
	orderID := r.FormValue("orderID")
	//查询数据库
	items, err := dao.GetOrderItemsByOrderID(orderID)
	if err != nil {
		log.Println(err)
		return
	}
	//解析模板
	t := template.Must(template.ParseFiles("src/Web/GoWeb/bookstore/views/pages/order/order_info.html"))
	_ = t.Execute(w, items)
}

// GetMyOrders 获取我的订单
func GetMyOrders(w http.ResponseWriter, r *http.Request) {
	flag, session := dao.IsLogin(r)
	if flag == true {
		//	获取用户id
		userID := session.UserID
		//获取用户订单信息
		orders, err := dao.GetMyOrdersByUserID(userID)
		if err != nil {
			log.Println(err)
			return
		}
		for _, v := range orders {
			v.UserName = session.UserName
			v.NoSend, v.NoSend, v.SendComplete = false, false, false
			switch v.State {
			case 0:
				v.NoSend = true
			case 1:
				v.SendComplete = true
			case 2:
				v.Complete = true
			}
		}
		t := template.Must(template.ParseFiles("src/Web/GoWeb/bookstore/views/pages/order/order.html"))
		_ = t.Execute(w, orders)
	} else {
		Login(w, r)
	}
}

// SendOrder 发货
func SendOrder(w http.ResponseWriter, r *http.Request) {
	order := r.FormValue("orderID")
	err := dao.UpDateOrderState(order, 1)
	if err != nil {
		log.Println(err)
		return
	}
	GetAllOrders(w, r)
}

// TakeOrder 收货
func TakeOrder(w http.ResponseWriter, r *http.Request) {
	order := r.FormValue("orderID")
	err := dao.UpDateOrderState(order, 2)
	if err != nil {
		log.Println(err)
		return
	}
	GetMyOrders(w, r)
}

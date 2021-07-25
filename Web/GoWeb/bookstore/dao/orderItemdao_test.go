package dao

import (
	"Web/GoWeb/bookstore/model"
	"fmt"
	"github.com/gofrs/uuid"
	"testing"
)

func TestOrder(t *testing.T) {
	fmt.Println("测试订单相关函数")
	t.Run("测试增加订单:", testAddOrder)
}

func testAddOrder(t *testing.T) {
	//生成订单号
	orderID, _ := uuid.NewV4()
	order := model.Order{
		OrderID:     orderID.String(),
		TotalCount:  2,
		TotalAmount: 400,
		State:       0,
		UserID:      1,
	}
	orderItem1 := model.OrderItem{
		Count:   1,
		Amount:  300,
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   300,
		ImgPath: "/static/img/default.jpg",
		OrderID: orderID.String(),
	}
	orderItem2 := model.OrderItem{
		Count:   1,
		Amount:  100,
		Title:   "红楼梦",
		Author:  "曹雪芹",
		Price:   100,
		ImgPath: "/static/img/default.jpg",
		OrderID: orderID.String(),
	}
	//保存订单
	AddOrder(&order)
	//保存订单项
	AddOrderItem(&orderItem1)
	AddOrderItem(&orderItem2)
}

func TestGetAllOrders(t *testing.T) {
	orders, err := GetAllOrders()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range orders {
		fmt.Println(v)
	}
}

func TestGetOrderItemsByOrderID(t *testing.T) {
	items, err := GetOrderItemsByOrderID("0f0b16a4-b7d8-4929-9c47-5f3c6796cf2b")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range items {
		fmt.Println(v)
	}
}

func TestGetMyOrdersByUserID(T *testing.T) {
	orders, err := GetMyOrdersByUserID(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range orders {
		fmt.Println(v)
	}
}

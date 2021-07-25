package dao

import (
	"Web/GoWeb/bookstore/model"
	"fmt"
	"github.com/gofrs/uuid"
	"testing"
	"time"
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
		CreateTime:  time.Now(),
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

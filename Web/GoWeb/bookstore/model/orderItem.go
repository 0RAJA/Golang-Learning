package model

// OrderItem 订单项
type OrderItem struct {
	OrderItemID string  //订单项ID
	Count       int     //订单项中图书数量
	Amount      float64 //订单项中图书金额
	Title       string  //图书书名
	Author      string  //作者
	Price       float64 //价格
	ImgPath     string  //图书封面
	OrderID     string  //订单项所属定单的编号
}

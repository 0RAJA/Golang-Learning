package model

// Order 订单
type Order struct {
	OrderID      string  //订单号
	CreateTime   string  //生成订单的时间
	TotalCount   int     //订单总数量
	TotalAmount  float64 //订单总金额
	State        int     //订单状态0未发货,1已发货,2交易完成
	UserID       int     //订单所属用户
	NoSend       bool    //state==0
	SendComplete bool    //state==1
	Complete     bool    //state==2
}

package model

type Session struct {
	SessionID string
	UserName  string
	UserID    int    //外键
	OrderID   string // 用来传给模板订单号
	IsRoot    bool
}

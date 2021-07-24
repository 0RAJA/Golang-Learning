package model

type Cart struct {
	CartID      string      //购物车ID
	CartItems   []*CartItem //购物车中所有的购物项
	TotalCount  int         //总数量
	TotalAmount float64     //总金额
	UserID      int         //当前购物车所属于的用户
}

func (cart *Cart) GetTotalCount() (ret int) {
	for _, v := range cart.CartItems {
		ret += v.Count
	}
	return
}

func (cart *Cart) GetTotalAmount() (ret float64) {
	for _, v := range cart.CartItems {
		ret += v.GetAmount()
	}
	return
}

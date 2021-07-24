package model

type CartItem struct {
	CartItemID int     //购物项ID
	Book       *Book   //图书信息
	Count      int     //图书数量
	Amount     float64 //图书总价格
	CartID     string  //当前购物车ID
}

func (CartItem *CartItem) GetAmount() float64 {
	return float64(CartItem.Count) * CartItem.Book.Price
}

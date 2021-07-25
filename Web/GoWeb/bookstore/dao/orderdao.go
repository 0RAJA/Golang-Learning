package dao

import (
	"Web/GoWeb/bookstore/model"
	"Web/GoWeb/bookstore/utils"
)

// AddOrder 填加订单
func AddOrder(order *model.Order) error {
	sql := "insert into orders (id, create_time, total_count, total_amount, state, user_id) values (?,?,?,?,?,?);"
	_, err := utils.DB.Exec(sql, order.OrderID, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteOrder(orderID string) error {
	err := DeleteOrderItem(orderID)
	if err != nil {
		return err
	}
	sql := "delete from orders where id = ?;"
	_, err = utils.DB.Exec(sql, orderID)
	if err != nil {
		return err
	}
	return nil
}

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

// GetAllOrders 获得所有订单
func GetAllOrders() (ret []*model.Order, err error) {
	sql := "select id, create_time, total_count, total_amount, state, user_id from orders;"
	rows, err := utils.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var order model.Order
		err = rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserID)
		if err != nil {
			return nil, err
		}
		ret = append(ret, &order)
	}
	return ret, nil
}

func GetMyOrdersByUserID(userID int) (ret []*model.Order, err error) {
	sql := "select id, create_time, total_count, total_amount, state, user_id from orders where user_id = ?;"
	rows, err := utils.DB.Query(sql, userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var order model.Order
		err := rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserID)
		if err != nil {
			return nil, err
		}
		ret = append(ret, &order)
	}
	return ret, err
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

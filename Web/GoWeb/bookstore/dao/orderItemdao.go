package dao

import (
	"Web/GoWeb/bookstore/model"
	"Web/GoWeb/bookstore/utils"
)

// AddOrderItem 插入订单项
func AddOrderItem(item *model.OrderItem) error {
	sql := "insert into order_items (count, amount, title, author, price, img_path, order_id) values (?,?,?,?,?,?,?);"
	_, err := utils.DB.Exec(sql, item.Count, item.Amount, item.Title, item.Author, item.Price, item.ImgPath, item.OrderID)
	if err != nil {
		return err
	}
	return nil
}

// GetOrderItemsByOrderID 通过订单号获取全部订单信息
func GetOrderItemsByOrderID(OrderID string) (ret []*model.OrderItem, err error) {
	sql := "select id, count, amount, title, author, price, img_path, order_id from order_items where order_id = ?;"
	rows, err := utils.DB.Query(sql, OrderID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var item model.OrderItem
		err := rows.Scan(&item.OrderItemID, &item.Count, &item.Amount, &item.Title, &item.Author, &item.Price, &item.ImgPath, &item.OrderID)
		if err != nil {
			return nil, err
		}
		ret = append(ret, &item)
	}
	return ret, nil
}

func DeleteOrderItem(orderID string) error {
	sql := "delete from order_items where order_id = ?;"
	_, err := utils.DB.Exec(sql, orderID)
	if err != nil {
		return err
	}
	return nil
}

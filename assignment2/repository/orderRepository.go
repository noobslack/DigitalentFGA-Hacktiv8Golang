package repository

import (
	"assignment2/model"
	"database/sql"
	"fmt"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (or *OrderRepository) CreateOrder(newOrder model.Order) (model.Order, error) {

	queryOrder := "INSERT INTO orders (customer_name) VALUES ($1) RETURNING order_id"

	var orderID int
	err := or.db.QueryRow(queryOrder, newOrder.Customer_name).Scan(&orderID)
	if err != nil {
		return model.Order{}, err
	}

	newOrder.Order_id = orderID

	for _, i := range newOrder.Item {
		queryItems := "INSERT INTO items (item_code, description, quantity, order_id) VALUES ($1, $2, $3, $4)"
		_, err := or.db.Exec(queryItems, i.Item_code, i.Description, i.Quantity, orderID)
		if err != nil {
			return model.Order{}, err
		}
	}
	return newOrder, nil
}

func (or *OrderRepository) GetOrder() ([]model.Order, error) {

	var orders = []model.Order{}

	query := `SELECT 
	orders.order_id, 
	orders.customer_name,
	orders.ordered_at,
	items.items_id,
	items.item_code,
	items.description,
	items.quantity, 
	items.order_id
	FROM orders orders JOIN items items ON 
	items.order_id = orders.order_id 
	WHERE items.order_id = orders.order_id`

	rows, err := or.db.Query(query)

	if err != nil {
		return orders, err
	}

	defer rows.Close()

	for rows.Next() {
		var o model.Order
		var i model.Item

		err := rows.Scan(
			&o.Order_id,
			&o.Customer_name,
			&o.Ordered_at,
			&i.Item_id,
			&i.Item_code,
			&i.Description,
			&i.Quantity,
			&i.Order_id)

		if err != nil {
			continue
		}
		var isExist bool
		for k := range orders {
			if orders[k].Order_id == o.Order_id {
				orders[k].Item = append(orders[k].Item, i)
				isExist = true
				break
			}
		}

		if !isExist {
			orders = append(orders, model.Order{
				Order_id:      o.Order_id,
				Customer_name: o.Customer_name,
				Ordered_at:    o.Ordered_at,
				Item:          []model.Item{i},
			})
		}
	}

	fmt.Println(orders)

	return orders, nil

}

func (or *OrderRepository) UpdateOrder(UpdateOrder model.Order) (model.Order, error) {

	queryOrder := "UPDATE orders SET customer_name = $1, ordered_at = $2 WHERE order_id = $3"

	_, err := or.db.Exec(queryOrder, UpdateOrder.Customer_name, UpdateOrder.Ordered_at, UpdateOrder.Order_id)
	if err != nil {
		return model.Order{}, err
	}

	for _, i := range UpdateOrder.Item {
		queryItems := "UPDATE items SET item_code = $1, description = $2, quantity =$3 WHERE items_id = $4 AND order_id=$5"
		_, err := or.db.Exec(queryItems, i.Item_code, i.Description, i.Quantity, i.Item_id, UpdateOrder.Order_id)
		if err != nil {
			return model.Order{}, err
		}
	}
	return UpdateOrder, nil

}

func (or *OrderRepository) DeleteOrder(id int) (model.Order, error) {

	var DeleteOrder model.Order

	queryOrder := "DELETE FROM orders where order_id=$1"

	_, err := or.db.Exec(queryOrder, id)

	if err != nil {
		return DeleteOrder, err
	}

	return DeleteOrder, nil

}

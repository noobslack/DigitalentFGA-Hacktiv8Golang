package repository

import "assignment2/model"

type IOrderRepository interface {
	CreateOrder(newOrder model.Order) (model.Order, error)
	GetOrder() ([]model.Order, error)
	UpdateOrder(UpdateOrder model.Order) (model.Order, error)
	DeleteOrder(id int) (model.Order, error)
}

package model

var Orders []Order = make([]Order, 0)

type Order struct {
	Order_id      int    `json:"-"`
	Customer_name string `json:"customer_name"`
	Ordered_at    string `json:"ordered_at"`
	Item          []Item `json:"items"`
}

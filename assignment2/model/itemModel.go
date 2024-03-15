package model

var Items []Item = make([]Item, 0)

type Item struct {
	Item_id     int    `json:"-"`
	Item_code   string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Order_id    int    `json:"-"`
}

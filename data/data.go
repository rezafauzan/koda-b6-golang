package data

type Menu struct {
	Id    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
}

type Cart struct {
	Id   int
	MenuData Menu
}

type CheckoutHistory struct {
	Id int
	CartData Cart
}

var MenuData []Menu
var CartData []Cart
var CheckoutHistoryData []CheckoutHistory

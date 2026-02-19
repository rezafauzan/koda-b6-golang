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

var MenuData []Menu
var CartData []Cart

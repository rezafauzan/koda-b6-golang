package lib

import (
	"fmt"
)

func ShowFood(menus []Menu, id int) {
	menu := menus[id]
	fmt.Printf("%d. %s Rp.%d,- \n", menu.Id, menu.Nama, menu.Harga)
}

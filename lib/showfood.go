package lib

import (
	"fmt"
	"rezafauzan/weekly-golang1/data"
)

func ShowFood(id int) {
	menu := data.MenuData[id]
	var navigator int
	fmt.Println("===================================================================")
	fmt.Println("==== (1). Untuk menambahkan ke keranjang                     ======")
	fmt.Println("==== (0). Untuk kembali ke list makanan                      ======")
	fmt.Println("===================================================================")
	fmt.Println("Input:")
	fmt.Scan(&navigator)
	if navigator == 0 {
		ShowMenu()
	} else if navigator == 1 {
		// cart = append(cart, menu)
	} else {
		ShowFood(navigator-1)
	}
	fmt.Printf("%d. %s Rp.%d,- \n", menu.Id, menu.Nama, menu.Harga)
}

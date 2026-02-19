package lib

import (
	"fmt"
	"rezafauzan/weekly-golang1/data"
)

func ShowFood(id int) {
	var menu data.Cart
	menu.MenuData = data.MenuData[id]
	menu.Id = 0
	var navigator int
	fmt.Println("===================================================================")
	fmt.Printf( "Product Id:%d \nProduct Name: %s \nProduct Price: Rp.%d,-        \n", menu.Id, menu.MenuData.Nama, menu.MenuData.Harga)
	fmt.Println("==== (1). Untuk menambahkan ke keranjang                     ======")
	fmt.Println("==== (0). Untuk kembali ke list makanan                      ======")
	fmt.Println("===================================================================")
	fmt.Println("Input:")
	fmt.Scan(&navigator)
	if navigator == 0 {
		ShowMenu()
	} else if navigator == 1 {
		data.CartData = append(data.CartData, menu)
		fmt.Printf("Product Id:%d \n Product Name: %s \n Product Price: Rp.%d,- ditambahkan ke keranjang \n", menu.Id, menu.MenuData.Nama, menu.MenuData.Harga)
		ShowMenu()
	}
}

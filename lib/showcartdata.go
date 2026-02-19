package lib

import (
	"fmt"
	"rezafauzan/weekly-golang1/data"
)

func ShowCartData() {
	var total int
	if len(data.CartData) > 0 {
		fmt.Println("===================================================================")
		for x := range data.CartData {
			total += data.CartData[x].MenuData.Harga			
			fmt.Printf("Item #%d \n Product Name: %s \n Product Price: Rp.%d,-\n", data.CartData[x].Id + 1, data.CartData[x].MenuData.Nama, data.CartData[x].MenuData.Harga)
		}
		fmt.Printf("Total: Rp.%d,-\n", total)
		} else {
			fmt.Println("==================== Keranjang Masih Kosong =======================")
		}
		// fmt.Println("===================================================================")
		// fmt.Println("==== (1). Untuk mengurangi item dari keranjang               ======")
		// fmt.Println("==== (0). Untuk kembali ke halaman utama                     ======")
		// fmt.Println("===================================================================")
		// fmt.Println("Input: ")
		// fmt.Scan(&navigator)
		// if navigator == 0 {
		// 	ShowMenuFoods()
		// } else if navigator == 2 {
		// 	ShowCartData()
		// }
	ShowMenu()
}

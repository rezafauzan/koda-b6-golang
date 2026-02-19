package lib

import (
	"fmt"
	"rezafauzan/weekly-golang1/data"
)

func ShowMenuFoods() {
	var navigator int
	for x := range data.MenuData {
		fmt.Printf("%d. %s Rp.%d,- \n", data.MenuData[x].Id, data.MenuData[x].Nama, data.MenuData[x].Harga)
	}
	fmt.Println("===================================================================")
	fmt.Println("==== (n).       Masukan angka untuk memilih makanan          ======")
	fmt.Println("==== (0).       Untuk kembali ke halaman utama               ======")
	fmt.Println("===================================================================")
	fmt.Println("Input: ")
	fmt.Scan(&navigator)
	if navigator <= len(data.MenuData) {
		if navigator == 0 {
			ShowMenu()
		} else {
			ShowFood(navigator - 1)
		}
	}
}

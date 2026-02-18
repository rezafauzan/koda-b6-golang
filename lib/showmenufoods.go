package lib

import (
	"fmt"
)

func ShowMenuFoods(menus []Menu) {
	var navigator int
	for x := range menus {
		fmt.Printf("%d. %s Rp.%d,- \n", menus[x].Id, menus[x].Nama, menus[x].Harga)
	}
	fmt.Println("===================================================================")
	fmt.Println("==== (n).       Masukan angka untuk memilih makanan          ======")
	fmt.Println("==== (0).       Untuk kembali ke halaman utama               ======")
	fmt.Println("===================================================================")
	fmt.Println("Input: ")
	fmt.Scanln(&navigator)
	if navigator <= len(menus) {
		if navigator == 0 {
			ShowMenu(menus)
		} else {
			ShowFood(menus, navigator - 1)
		}
	}
}

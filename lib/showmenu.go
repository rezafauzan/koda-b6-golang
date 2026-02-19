package lib

import (
	"fmt"
)

func ShowMenu() {
	var navigator int
	fmt.Println("===================================================================")
	fmt.Println("==== Selamat Datang di Program Interactive Pembelian Makanan ======")
	fmt.Println("==== (1) Lihat list menu makanan                             ======")
	fmt.Println("==== (2) Lihat keranjang                                     ======")
	fmt.Println("==== (3) Checkout keranjang                                  ======")
	fmt.Println("==== (4) Lihat History Invoice Terbuat                       ======")
	fmt.Println("===================================================================")
	fmt.Println("====         Silahkan pilih menu dengan memasukan angka      ======")
	fmt.Println("===================================================================")
	fmt.Println("Input: ")
	fmt.Scan(&navigator)
	if navigator == 1 {
		ShowMenuFoods()
	}
}

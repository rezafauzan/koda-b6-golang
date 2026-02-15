package main

import (
	"fmt"
	"rezafauzan/weekly-golang1/lib"
)

var menus []lib.Menu

func ShowMenuFoods(menus []lib.Menu) {
	var navigator int
	for x:= range menus{
		fmt.Printf("%d. %s Rp.%d,- \n", menus[x].Id, menus[x].Nama, menus[x].Harga)
	}
	fmt.Println("===================================================================")
	fmt.Println("==== (n).       Masukan angka untuk memilih makanan          ======")
	fmt.Println("==== ('kata').  Masukan kata untuk mencari menu              ======")
	fmt.Println("==== (0).       Untuk kembali ke halaman utama               ======")
	fmt.Println("===================================================================")
	fmt.Println("Input: ")
	fmt.Scanln(&navigator)
	if(navigator == 1){
		ShowMenuFoods(menus)
	}
}

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
	fmt.Scanln(&navigator)
	if(navigator == 1){
		ShowMenuFoods(menus)
	}
}

func main() {
	var URL string = "https://raw.githubusercontent.com/rezafauzan/koda-b6-golang/refs/heads/main/assets/data/menu.json"
	lib.Fetch(URL, &menus)
	if len(menus) > 0 {
		ShowMenu()
	}
}

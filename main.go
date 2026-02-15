package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type menu struct {
	Id    int `json:id`
	Nama  string `json:nama`
	Harga int `json:harga`
}

var menus []menu

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
}

func main() {
	var URL string = "https://raw.githubusercontent.com/rezafauzan/koda-b6-golang/refs/heads/main/assets/data/menu.json"
	response, errors := http.Get(URL)
	if errors != nil {
		fmt.Println(errors)
		} else {
			data, errors := io.ReadAll(response.Body)
			if errors != nil{
			fmt.Println(errors)
			}else{
				errors := json.Unmarshal(data, &menus)
				if errors != nil {
					fmt.Println(errors)
				} else {
					fmt.Println(menus)
					ShowMenu()
				}
		}
	}
}

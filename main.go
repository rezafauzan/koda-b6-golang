package main

import (
	"rezafauzan/weekly-golang1/data"
	"rezafauzan/weekly-golang1/lib"
)

func main() {
	var URL string = "https://raw.githubusercontent.com/rezafauzan/koda-b6-golang/refs/heads/main/assets/data/menu.json"
	lib.Fetch(URL, &data.MenuData)
	if len(data.MenuData) > 0 {
		lib.ShowMenu()
	}
}

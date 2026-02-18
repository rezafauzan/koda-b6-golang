package main

import (
	"rezafauzan/weekly-golang1/lib"
)

var menus []lib.Menu

func main() {
	var URL string = "https://raw.githubusercontent.com/rezafauzan/koda-b6-golang/refs/heads/main/assets/data/menu.json"
	lib.Fetch(URL, &menus)
	if len(menus) > 0 {
		lib.ShowMenu(menus)
	}
}

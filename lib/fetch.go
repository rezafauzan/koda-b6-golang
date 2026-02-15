package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Menu struct {
	Id    int `json:"id"`
	Nama  string `json:"nama"`
	Harga int `json:"harga"`
}

func Fetch(URL string, menus *[]Menu){
	response, errors := http.Get(URL)
	if errors != nil {
		fmt.Println(errors)
	} else {
		data, errors := io.ReadAll(response.Body)
		if errors != nil {
			fmt.Println(errors)
		} else {
			errors := json.Unmarshal(data, &menus)
			if errors != nil {
				fmt.Println(errors)
			} else {
				return
			}
		}
	}
}

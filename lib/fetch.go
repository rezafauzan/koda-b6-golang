package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"rezafauzan/weekly-golang1/data"
)

func Fetch(URL string, MenuData *[]data.Menu) {
	response, errors := http.Get(URL)
	if errors != nil {
		fmt.Println(errors)
	} else {
		data, errors := io.ReadAll(response.Body)
		if errors != nil {
			fmt.Println(errors)
		} else {
			errors := json.Unmarshal(data, &MenuData)
			if errors != nil {
				fmt.Println(errors)
			} else {
				return
			}
		}
	}
}

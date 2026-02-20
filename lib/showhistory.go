package lib

import (
	"fmt"
	"rezafauzan/weekly-golang1/data"
)

func ShowHistory() {
	if len(data.CheckoutHistoryData) > 0 {
		fmt.Println("===================================================================")
		for x := range data.CheckoutHistoryData {
			fmt.Printf("History #%d \n", data.CheckoutHistoryData[x].Id+1)
			for y := range data.CheckoutHistoryData[x].CartDataHistory {
				fmt.Printf("Item #%d \n Product Name: %s \n Product Price: Rp.%d,-\n", data.CheckoutHistoryData[x].CartDataHistory[y].Id+1, data.CheckoutHistoryData[x].CartDataHistory[y].MenuData.Nama, data.CheckoutHistoryData[x].CartDataHistory[y].MenuData.Harga)
			}
		}
		fmt.Println("===================================================================")
	} else {
		fmt.Println("==================== Riwayat Belanja Masih Kosong =======================")
	}
	ShowMenu()
}

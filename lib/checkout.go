package lib

import (
	"fmt"
	"rezafauzan/weekly-golang1/data"
)

func Checkout() {
	var total int
	if len(data.CartData) > 0 {
		fmt.Println("")
		fmt.Println("===================================================================")
		fmt.Println("================== Pembayaran Berhasil ============================")
		fmt.Println("===================================================================")
		fmt.Println("")
		for x := range data.CartData {
			total += data.CartData[x].MenuData.Harga			
			fmt.Printf("Item #%d \n Product Name: %s \n Product Price: Rp.%d,-\n", data.CartData[x].Id + 1, data.CartData[x].MenuData.Nama, data.CartData[x].MenuData.Harga)
		}
		fmt.Printf("Total: Rp.%d,-\n", total)
		checkoutData := data.CheckoutHistory{
			Id: len(data.CheckoutHistoryData),
			CartDataHistory: data.CartData,
		}
		data.CheckoutHistoryData = append(data.CheckoutHistoryData, checkoutData)
		data.CartData = []data.Cart{}
		fmt.Println("===================================================================")
		fmt.Println("===================================================================")
		fmt.Println("")
		} else {
			fmt.Println("===================================================================")
			fmt.Println("==================== Keranjang Masih Kosong =======================")
			fmt.Println("===================================================================")
			fmt.Println("")
		}
	ShowMenu()
}

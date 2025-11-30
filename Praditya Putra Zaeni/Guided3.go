package main

import "fmt"

func main() {
		var kendaraan string
		var durasi int
		var biaya int

		fmt.Print(" jenis kendaraan (Motor/Mobil/Truk): ")
		fmt. Scan(&kendaraan)
		fmt.Print("durasi parkir : ")
		fmt. Scan(&durasi)

	switch {
		case kendaraan == "Motor" && durasi >= 1 && durasi <= 2: 
			biaya = 7000
		case kendaraan == "Motor" && durasi > 2:
			biaya = 9000
		case kendaraan == "Mobil" && durasi >= 1 && durasi <= 2:
			biaya = 15000
		case kendaraan == "Mobil" && durasi > 2:
			biaya = 20000
		case kendaraan == "Truk" && durasi >= 1 && durasi <= 2:
			biaya = 25000
		case kendaraan == "Truk" && durasi > 2:
			biaya = 35000
	default:
		fmt.Println("Data Tidak valid")
	}
		fmt.Printf("Tarif Parkir: Rp %d\n", biaya)
}
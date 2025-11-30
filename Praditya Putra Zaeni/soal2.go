package main

import "fmt"

func main() {
		var kendaraan string
		var durasi int
		var tarif int

		fmt.Print(" jenis kendaraan (Motor/Mobil/Truk): ")
		fmt. Scan(&kendaraan)
		fmt.Print("durasi parkir : ")
		fmt. Scan(&durasi)

	switch {
		case kendaraan == "motor" && durasi >= 1 && durasi <= 2: 
			tarif = 2000
		case kendaraan == "motor" && durasi > 2:
			tarif = 6000
		case kendaraan == "mobil" && durasi >= 1 && durasi <= 2:
			tarif = 5000
		case kendaraan == "mobil" && durasi > 2:
			tarif = 5000
		case kendaraan == "truk" && durasi >= 1 && durasi <= 2:
			tarif = 8000
		case kendaraan == "truk" && durasi > 2:
			tarif = 40000
	default:
		fmt.Println("Data Tidak valid")
	}
		fmt.Printf("Tarif Parkir: Rp %d\n", tarif)
}
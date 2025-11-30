package main

import (
	"fmt"
)

func main() {
	var jenis string
	var durasi float64
	var tarif int

	fmt.Print("Jenis kendaraan (motor/mobil/truk): ")
	fmt.Scan(&jenis)
	fmt.Print("Durasi parkir (jam, cth 3.5): ")
	fmt.Scan(&durasi)

	switch jenis {
	case "motor":
		tarif = 2000
	case "mobil":
		tarif = 5000
	case "truk":
		tarif = 8000
	default:
		tarif = 0
	}
	if tarif > 0 {
		if durasi < 1.0 {
			durasi = 1.0
		}
		jamDikenakanBiaya := int(durasi)
		if durasi > float64(jamDikenakanBiaya) {
			jamDikenakanBiaya++
		}
		totalBiaya := jamDikenakanBiaya * tarif
		fmt.Printf("\nTotal Biaya Parkir: Rp %d\n", totalBiaya)
	} else {
	}
}

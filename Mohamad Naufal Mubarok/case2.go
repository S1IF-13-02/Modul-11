package main

import (
	"fmt"
)

func main() {
	var jenis string
	var durasi int

	fmt.Print("Masukkan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scan(&jenis)

	fmt.Print("Masukkan durasi parkir (jam): ")
	fmt.Scan(&durasi)

	if durasi < 1 {
		durasi = 1
	}

	var tarif int

	switch jenis {
	case "motor":
		tarif = 2000
	case "mobil":
		tarif = 5000
	case "truk":
		tarif = 8000
	default:
		fmt.Println("Jenis kendaraan tidak valid.")
		return
	}

	total := tarif * durasi
	fmt.Println("Total biaya parkir: Rp", total)
}
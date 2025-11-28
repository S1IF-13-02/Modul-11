package main

import "fmt"

func main() {
	var jenis string
	var jam int

	fmt.Print("Masukkan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scan(&jenis)

	fmt.Print("Masukkan durasi parkir (jam): ")
	fmt.Scan(&jam)

	if jam < 1 {
		jam = 1
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

	total := tarif * jam

	fmt.Printf("Total biaya parkir: Rp %d\n", total)
}

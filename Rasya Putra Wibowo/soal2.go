package main

import "fmt"

func main() {
	var kendaraan string
	var jam int

	fmt.Print("Masukkan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukkan durasi parkir (jam): ")
	fmt.Scan(&jam)

	if jam < 1 {
		jam = 1
	}

	var tarif int

	switch kendaraan {
	case "motor":
		tarif = 2000
	case "mobil":
		tarif = 5000
	case "truk":
		tarif = 8000
	default:
		fmt.Println("Jenis kendaraan tidak valid")
		return
	}

	total := tarif * jam
	fmt.Println("Total biaya parkir: Rp", total)
}

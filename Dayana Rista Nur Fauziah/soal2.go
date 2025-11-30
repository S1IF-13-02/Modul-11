package main

import "fmt"

func main() {
	var jenis string
	var jam int

	fmt.Print("Jenis kendaraan (motor/mobil/truk): ")
	fmt.Scan(&jenis)

	fmt.Print("Durasi parkir (jam): ")
	fmt.Scan(&jam)

	if jam < 1 {
		jam = 1
	}

	tarif := 0

	switch jenis {
	case "motor":
		tarif = 2000 * jam
	case "mobil":
		tarif = 5000 * jam
	case "truk":
		tarif = 8000 * jam
	default:
		fmt.Println("Jenis kendaraan tidak valid.")
		return
	}

	fmt.Println("Total biaya parkir:", tarif)
}

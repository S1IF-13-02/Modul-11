package main

import "fmt"

func main() {
	var jenis string
	var durasi int
	var tarif int

	fmt.Print("Masukkan jenis kendaraan (Motor/Mobil/Truk): ")
	fmt.Scan(&jenis)

	fmt.Print("Masukkan durasi parkir (jam): ")
	fmt.Scan(&durasi)

	switch jenis {
	case "Motor":
		if durasi >= 1 && durasi <= 2 {
			tarif = 7000
		} else if durasi > 2 {
			tarif = 9000
		} else {
			tarif = 0
		}

	case "Mobil":
		if durasi >= 1 && durasi <= 2 {
			tarif = 15000
		} else if durasi > 2 {
			tarif = 20000
		} else {
			tarif = 0
		}

	case "Truk":
		if durasi >= 1 && durasi <= 2 {
			tarif = 25000
		} else if durasi > 2 {
			tarif = 35000
		} else {
			tarif = 0
		}

	default:
		fmt.Println("Jenis kendaraan tidak valid.")
		return
	}

	fmt.Println("Tarif Parkir: Rp", tarif)
}

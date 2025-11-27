package main

import "fmt"

func main() {
	var kendaraan string
	var durasi, tarif float64
	fmt.Print("Masukkan jenis kendaraan (mobil/motor/truk): ")
	fmt.Scanln(&kendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scanln(&durasi)

	switch kendaraan {
	case "motor":
		if durasi > 1 {
			tarif = durasi * 2000
			fmt.Print("tarif parkir Rp.", tarif)
		} else if durasi > 0 && durasi <= 1 {
			tarif = 2000
			fmt.Print("tarif parkir Rp.", tarif)
		} else {
			fmt.Println("Jenis kendaraan / durasi tidak valid.")
		}
	case "mobil":
		if durasi > 1 {
			tarif = durasi * 5000
			fmt.Print("tarif parkir Rp.", tarif)
		} else if durasi > 0 && durasi <= 1 {
			tarif = 5000
			fmt.Print("tarif parkir Rp.", tarif)
		} else {
			fmt.Println("Jenis kendaraan / durasi tidak valid.")
		}
	case "truk":
		if durasi > 1 {
			tarif = durasi * 8000
			fmt.Print("tarif parkir Rp.", tarif)
		} else if durasi > 0 && durasi <= 1 {
			tarif = 8000
			fmt.Print("tarif parkir Rp.", tarif)
		} else {
			fmt.Println("Jenis kendaraan / durasi tidak valid.")
		}
	default:
		fmt.Println("Jenis kendaraan / durasi tidak valid.")

	}

}

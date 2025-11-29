package main

import (
	"fmt"
	"strings"
)

func main() {
	var kendaraan string
	var durasi int

	fmt.Print("Masukkan jenis kendaraan (Motor/Mobil/Truk): ")
	fmt.Scan(&kendaraan)

	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	kendaraan = strings.ToLower(kendaraan)

	switch kendaraan {
	case "motor":
		switch {
		case durasi >= 1 && durasi <= 2:
			fmt.Println("Tarif Parkir: Rp 7000")
		case durasi > 2:
			fmt.Println("Tarif Parkir: Rp 9000")
		default:
			fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
			fmt.Println("Tarif Parkir: Rp 0")
		}

	case "mobil":
		switch {
		case durasi >= 1 && durasi <= 2:
			fmt.Println("Tarif Parkir: Rp 15.000")
		case durasi > 2:
			fmt.Println("Tarif Parkir: Rp 20.000")
		default:
			fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
			fmt.Println("Tarif Parkir: Rp 0")
		}

	case "truk":
		switch {
		case durasi >= 1 && durasi <= 2:
			fmt.Println("Tarif Parkir: Rp 25.000")
		case durasi > 2:
			fmt.Println("Tarif Parkir: Rp 30.000")
		default:
			fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
			fmt.Println("Tarif Parkir: Rp 0")
		}

	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
		fmt.Println("Tarif Parkir: Rp 0")
	}
}

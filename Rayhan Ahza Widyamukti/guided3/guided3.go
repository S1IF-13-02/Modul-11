package main

import "fmt"

func main() {
	var kendaraan string
	var durasi int

	fmt.Println("Masukan Jenis Kendaraan (Motor/Mobil/Truk): ")
	fmt.Scan(&kendaraan)

	fmt.Println("Masukan Durasi Parkir: ")
	fmt.Scan(&durasi)

	switch kendaraan {
	case "Motor":
		if durasi <= 2 {
			fmt.Println("Tarif Parkir: Rp.7000")
		} else {
			fmt.Println("Tarif Parkir: Rp9.000")
		}
	case "Mobil":
		if durasi <= 2 {
			fmt.Println("Tarif Parkir: Rp15.000")
		} else {
			fmt.Println("Tarif Parkir: Rp20.000")
		}
	case "Truk":
		if durasi <= 2 {
			fmt.Println("Tarif Parkir: 25.000")
		} else {
			fmt.Println("Tarif Parkir: 35.000")
		}

	default:
		fmt.Println("Jenis Kendaraan atau Durasi Parkir Tidak Valid")
	}

}

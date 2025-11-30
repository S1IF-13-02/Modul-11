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
		if durasi <= 1 {
			fmt.Println("Tarif Parkir: Rp.2000")
		} else {
			fmt.Println("Tarif Parkir: Rp", 2000*durasi)
		}
	case "Mobil":
		if durasi <= 1 {
			fmt.Println("Tarif Parkir: Rp.5000")
		} else {
			fmt.Println("Tarif Parkir: Rp", 5000*durasi)
		}
	case "Truk":
		if durasi <= 1 {
			fmt.Println("Tarif Parkir: Rp.8000")
		} else {
			fmt.Println("Tarif Parkir: Rp", 8000*durasi)
		}

	default:
		fmt.Println("Jenis Kendaraan atau Durasi Parkir Tidak Valid")
	}

}

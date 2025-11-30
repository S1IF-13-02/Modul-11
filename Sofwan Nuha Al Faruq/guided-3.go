package main

import "fmt"

func main() {
	var jenisKendaraan string
	var lamaParkir int
	var tarif int

	fmt.Print("Masukkan jenis kendaraan (Motor/Mobil/Truk): ")
	fmt.Scanln(&jenisKendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scanln(&lamaParkir)

	if lamaParkir <= 0 {
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
		fmt.Println("Tarif Parkir: Rp 0")
		return
	}

	switch jenisKendaraan {
	case "Motor":
		if lamaParkir <= 2 {
			tarif = 7000
		} else {
			tarif = 9000
		}
		fmt.Printf("Tarif Parkir: Rp %d\n", tarif)
	case "Mobil":
		if lamaParkir <= 2 {
			tarif = 15000
		} else {
			tarif = 20000
		}
		fmt.Printf("Tarif Parkir: Rp %d\n", tarif)
	case "Truk":
		if lamaParkir <= 2 {
			tarif = 25000
		} else {
			tarif = 35000
		}
		fmt.Printf("Tarif Parkir: Rp %d\n", tarif)
	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
		fmt.Println("Tarif Parkir: Rp 0")
	}
}

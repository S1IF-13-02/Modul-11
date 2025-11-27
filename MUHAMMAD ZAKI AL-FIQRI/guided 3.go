package main

import "fmt"

func main() {
	var kendaraan string
	var durasi int
	fmt.Print("Masukkan jenis kendaraan (mobil/motor/truk): ")
	fmt.Scanln(&kendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scanln(&durasi)

	switch kendaraan {
	case "motor":
		if durasi >= 1 && durasi <= 2 {
			fmt.Println("tarif parkir : Rp 7000")
		} else {
			fmt.Println("tarif parkir : Rp 9000")
		}
	case "mobil":
		if durasi >= 1 && durasi <= 2 {
			fmt.Println("tarif parkir : Rp 15000")
		} else {
			fmt.Println("tarif parkir : Rp 20000")
		}
	case "truk":
		if durasi >= 1 && durasi <= 2 {
			fmt.Println("tarif parkir : Rp 25000")
		} else {
			fmt.Println("tarif parkir : Rp 35000")
		}
	default:
		fmt.Println("Jenis kendaraan / durasi tidak valid.")
		fmt.Println("tarif parkir : Rp 0")
	}

}

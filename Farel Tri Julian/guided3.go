package main

import "fmt"

func main() {
	var kendaraan string
	var jam int
	fmt.Print("masukan jenis kendaraan (mototr/mobil/truk): ")
	fmt.Scan(&kendaraan)
	fmt.Print("masukan durasi parkir (dalam jam): ")
	fmt.Scan(&jam)

	switch kendaraan {
	case "motor":
		if jam <= 2 && jam >= 1 {
			fmt.Println("Tarif parkir : Rp 7.000")
		} else {
			fmt.Println("Tarif parkir : Rp 9.000")
		}
	case "mobil":
		if jam <= 2 && jam >= 1 {
			fmt.Println("Tarif parkir : Rp 15.000")
		} else {
			fmt.Println("Tarif parkir : Rp 20.000")
		}
	case "truk":
		if jam <= 2 && jam >= 1 {
			fmt.Println("Tarif parkir : Rp 25.000")
		} else {
			fmt.Println("Tarif parkir : Rp 35.000")
		}
	default:
		fmt.Println("jenis kendaraan atau durasi tidak valid")
		fmt.Println("tarif parkir Rp 0")
	}
}

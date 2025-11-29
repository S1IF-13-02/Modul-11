package main

import "fmt"

func main() {
	var kendaraan, parkir string
	var jam int

	fmt.Print("masukan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scan(&kendaraan)

	fmt.Print("masukan durasi parkir(dalam jam): ")
	fmt.Scan(&jam)

	switch kendaraan {
	case "motor":
		if jam > 2 {
			parkir = "9000"
		} else {
			parkir = "7000"
		}
	case "mobil":
		if jam > 2 {
			parkir = "20000"
		} else {
			parkir = "15000"
		}
	case "truk":
		if jam > 2 {
			parkir = "35000"
		} else {
			parkir = "25000"
		}
	default:
		fmt.Println("jenis kendaraan tidak valid")
	}
	fmt.Println("Tarif parkir : ", parkir)
}

package main

import "fmt"
func main() {
	var kendaraan, parkir string
	var jam int

	fmt.Print("Masukan jenis kendaraan (motor/mobil/truk) : ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukan Durasi Parkir (dalam jam) : ")
	fmt.Scan(&jam)

	switch kendaraan {
	case "motor":
		if jam > 2 {
			parkir = "9.000"
		} else {
			parkir = "7.000"
		}
	case "mobil":
		if jam > 2 {
			parkir = "20.000"
		} else {
			parkir = "15.000"
		}
	case "truk":
		if jam > 2 {
			parkir = "35.000"
		} else {
			parkir = "25.000"
		}
	default:
		fmt.Println("Jenis Kendaraan Tidak Valid")
		return
	}

	fmt.Printf("Tarif Parkir: Rp %s\n", parkir)
}

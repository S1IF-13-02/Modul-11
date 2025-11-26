package main

import "fmt"
func main() {
	var kendaraan string
	var jam int
	var tarif int

	fmt.Print("Masukan jenis kendaraan (motor/mobil/truk) : ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukan Durasi Parkir (dalam jam) : ")
	fmt.Scan(&jam)

	if jam < 1 {
		jam = 1
	}

	switch kendaraan {
	case "motor":
		tarif = 2000
	case "mobil":
		tarif = 5000
	case "truk":
		tarif = 8000
	default:
		fmt.Println("Jenis Kendaraan Tidak Valid")
		return
	}

	total := tarif * jam

	fmt.Printf("Total Tarif Parkir: Rp %d\n", total)
}

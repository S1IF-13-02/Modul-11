package main

import "fmt"

func main() {
	var kendaraan string
	var durasi, tarif int
	fmt.Print("Jenis Kendaraan: ")
	fmt.Scan(&kendaraan)
	fmt.Print("Durasi Parkir: ")
	fmt.Scan(&durasi)
	if durasi < 1 {
		durasi = 1
	}
	switch kendaraan {
	case "Motor":
		tarif = 2000 * durasi
	case "Mobil":
		tarif = 5000 * durasi
	case "Truk":
		tarif = 8000 * durasi
	default:
		fmt.Println("Jenis Kendaraan dan Durasi Parkir Tidak Valid")
	}
	fmt.Printf("Total Biaya Parkir: %d\n", tarif)
}

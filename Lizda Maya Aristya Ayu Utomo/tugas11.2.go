package main

import "fmt"

func main() {
	var durasi int
	var kendaraan string

	fmt.Scan(&kendaraan, &durasi)

	if durasi < 1 {
		durasi = 1
	}
	var tarif int

	switch kendaraan {
	case "motor":
		tarif = 2000
	case "mobil":
		tarif = 5000
	case "truk":
		tarif = 8000
	default:
		fmt.Println("jenis kendaraan tidak valid")
		return

	}
	
	total := tarif * durasi
	fmt.Printf("total biaya parkir: Rp %d\n", total)

}

package main

import "fmt"

func main() {
	var kendaraan string
	var durasi int
	var total int
	var tarif int

	fmt.Print("Masukkan jenis kendaraan (Motor/Mobil/Truk): ") 
	fmt.Scan(&kendaraan) 
    fmt.Print("Masukkan durasi parkir (dalam jam): ") 
    fmt.Scan(&durasi) 

	switch kendaraan {
	case "Motor":
		if durasi > 1 {
			tarif = 2000
		}else {
			tarif = 2000
		}
	case "Mobil":
		if durasi > 1 {
			tarif = 5000
		}else {
			tarif = 5000
		}
	case "Truk":
		if durasi > 1 {
			tarif = 8000
		}else {
			tarif = 8000
		}
	default:
		fmt.Println("Jenis kendaraan tidak valid")
	}
	total = tarif * durasi
	fmt.Println("Tarif parkir = ", total)
}
package main

import "fmt"

func main() {
	var Kendaraan string
	var durasi int
	fmt.Print("Masukkan jenis kendaraan (mobil/motor/truk): ")
	fmt.Scan(&Kendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	var tarifPer1sampai2Jam int
	switch Kendaraan {
	case "motor":
		tarifPer1sampai2Jam = 7000
		tarifLebihDari2Jam := 9000
		if durasi > 2 {
			total := tarifLebihDari2Jam
			fmt.Printf("Tarif Parkir: Rp.%d", total)
		} else {
			total := tarifPer1sampai2Jam
			fmt.Printf("Tarif Parkir: Rp.%d", total)
		}
	case "mobil":
		tarifPer1sampai2Jam = 15000
		tarifLebihDari2Jam := 20000
		if durasi > 2 {
			total := tarifLebihDari2Jam
			fmt.Printf("Tarif Parkir: Rp.%d", total)
		} else {
			total := tarifPer1sampai2Jam
			fmt.Printf("Tarif Parkir: Rp.%d", total)
		}
	case "truk":
		tarifPer1sampai2Jam = 25000
		tarifLebihDari2Jam := 35000
		if durasi > 2 {
			total := tarifLebihDari2Jam
			fmt.Printf("Tarif Parkir: Rp.%d", total)
		} else {
			total := tarifPer1sampai2Jam
			fmt.Printf("Tarif Parkir: Rp.%d", total)
		}
	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid.")
		fmt.Println("Tarif Parkir: Rp.0")
	}
}

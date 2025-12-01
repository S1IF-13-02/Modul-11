package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputKendaraan string
	var durasiParkir int

	fmt.Print("Masukkan jenis kendaraan (Motor/Mobil/Truk): ")
	fmt.Scan(&inputKendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasiParkir)

	kendaraan := strings.ToLower(inputKendaraan)
	tarifParkir := 0

	switch kendaraan {
	case "motor":
		if durasiParkir < 1 {
			tarifParkir = 2000
			break
		}

		tarifParkir = 2000 * durasiParkir
	case "mobil":
		if durasiParkir < 1 {
			tarifParkir = 5000
			break
		}

		tarifParkir = 5000 * durasiParkir
	case "truk":
		if durasiParkir < 1 {
			tarifParkir = 8000
			break
		}

		tarifParkir = 8000 * durasiParkir
	}

	fmt.Printf("Tarif Parkir: Rp. %d\n", tarifParkir)
}

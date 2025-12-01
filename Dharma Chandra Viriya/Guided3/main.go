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
		if durasiParkir > 2 {
			tarifParkir = 9000
			break
		}

		tarifParkir = 7000
	case "mobil":
		if durasiParkir > 2 {
			tarifParkir = 20000
			break
		}

		tarifParkir = 15000
	case "truk":
		if durasiParkir > 2 {
			tarifParkir = 35000
			break
		}

		tarifParkir = 25000
	}

	fmt.Printf("Tarif Parkir: Rp. %d\n", tarifParkir)
}

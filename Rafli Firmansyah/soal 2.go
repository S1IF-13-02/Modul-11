package main

import (
	"fmt"
	"strings"
)

func main() {
	var kendaraan string
	var durasi, tarifPerJam, total int

	fmt.Print("Masukkan jenis kendaraan (motor/mobil/truk): ")
	fmt.Scan(&kendaraan)

	fmt.Print("Masukkan durasi parkir (jam): ")
	fmt.Scan(&durasi)

	if durasi < 1 {
		durasi = 1
	}

	kendaraan = strings.ToLower(kendaraan)

	switch kendaraan {
	case "motor":
		tarifPerJam = 2000
	case "mobil":
		tarifPerJam = 5000
	case "truk":
		tarifPerJam = 8000
	default:
		fmt.Println("Jenis kendaraan tidak dikenali.")
		return
	}

	total = tarifPerJam * durasi

	fmt.Println("Total biaya parkir: Rp", total)
}

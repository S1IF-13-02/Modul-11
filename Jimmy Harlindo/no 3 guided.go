package main

import (
	"fmt"
	"strings"
)

func main() {
	var jenis string
	var durasi int

	fmt.Print("Masukkan jenis kendaraan (Motor/Mobil/Truk): ")
	fmt.Scanln(&jenis)

	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scanln(&durasi)

	j := strings.ToLower(jenis)
	tarif := 0

	switch j {
	case "motor":
		if durasi >= 1 && durasi <= 2 {
			tarif = 7000
		} else if durasi > 2 {
			tarif = 9000
		}

	case "mobil":
		if durasi >= 1 && durasi <= 2 {
			tarif = 15000
		} else if durasi > 2 {
			tarif = 20000
		}

	case "truk":
		if durasi >= 1 && durasi <= 2 {
			tarif = 25000
		} else if durasi > 2 {
			tarif = 35000
		}

	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
		fmt.Println("Tarif Parkir: Rp 0")
		return
	}

	fmt.Println("Tarif Parkir: Rp", tarif)
}

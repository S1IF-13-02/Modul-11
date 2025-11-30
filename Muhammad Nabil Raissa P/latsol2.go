package main

import "fmt"

func main() {
	var tipeKendaraan string
	var durasiJam int
	var biayaDasar int

	fmt.Print("Masukan jenis kendaraan (motor/mobil/truk) : ")
	fmt.Scan(&tipeKendaraan)
	fmt.Print("Masukan Durasi Parkir (dalam jam) : ")
	fmt.Scan(&durasiJam)

	if durasiJam < 1 {
		durasiJam = 1
	}

	switch tipeKendaraan {
	case "motor":
		biayaDasar = 2000
	case "mobil":
		biayaDasar = 5000
	case "truk":
		biayaDasar = 8000
	default:
		fmt.Println("Jenis Kendaraan Tidak Valid")
		return
	}

	totalBiaya := biayaDasar * durasiJam

	fmt.Printf("Total Tarif Parkir: Rp %d\n", totalBiaya)
}

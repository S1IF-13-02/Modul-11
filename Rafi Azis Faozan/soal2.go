package main

import "fmt"

func main() {
	var kendaraan string
	var durasi int
	fmt.Print("Masukkan jenis kendaraan: ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukkan durasi parkir: ")
	fmt.Scan(&durasi)
	switch kendaraan {
	case "motor":
		biaya := 2000 * durasi
		fmt.Print("Rp. ", biaya)
	case "mobil":
		biaya := 5000 * durasi
		fmt.Print("Rp. ", biaya)
	case "truk":
		biaya := 8000 * durasi
		fmt.Print("Rp. ", biaya)
	}
}

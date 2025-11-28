package main

import (
	"fmt"
	"strings"
)

func main() {
	
	var jenis_kendaraan string
	var durasi_parkir, total_biaya int

	fmt.Print("Masukkan jenis kendaraan (Motor/Mobil/Truk): ")
	fmt.Scan(&jenis_kendaraan)

	jenis_kendaraan = strings.ToLower(jenis_kendaraan)

	fmt.Print("Durasi parkir (jam): ")
	fmt.Scan(&durasi_parkir)

	if durasi_parkir < 0{
		fmt.Println("Durasi parkir tidak valid")
		return
	}

	if durasi_parkir < 1{
		durasi_parkir = 1
	}

	switch {
	case "motor" == jenis_kendaraan:
		total_biaya = 2000 * durasi_parkir
		fmt.Printf("Total biaya parkir motor anda adalah: Rp.%d\n",total_biaya)
	case "mobil" == jenis_kendaraan:
		total_biaya = 5000 * durasi_parkir
		fmt.Printf("Total biaya parkir mobil anda adalah: Rp.%d\n",total_biaya)
	case "truk" == jenis_kendaraan:
		total_biaya = 8000 * durasi_parkir
		fmt.Printf("Total biaya parkir truk anda adalah: Rp.%d\n",total_biaya)
	default:
		fmt.Println("Jenis kendaraan tidak dikenali")
	}
}

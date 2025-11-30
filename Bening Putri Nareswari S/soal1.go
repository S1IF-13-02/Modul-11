package main

import (
	"fmt"
)

func main() {
	var ph float64
	var hasil string

	fmt.Print("Masukkan kadar pH air (0.0 - 14.0): ")
	fmt.Scan(&ph)
	switch {
	case ph < 0.0 || ph > 14.0:
		hasil = "Input tidak valid, rentang pH harus antara 0.0 dan 14.0."
	case ph >= 6.5 && ph <= 8.6:
		hasil = "Air Layak Minum (pH >= 6.5 dan pH <= 8.6)"
	default:
		hasil = "Air Tidak Layak Minum (pH < 6.5 dan pH > 8.6)"
	}

	fmt.Println(hasil)
}

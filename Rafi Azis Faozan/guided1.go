package main

import "fmt"

func main() {
	var jam int
	var waktu string
	fmt.Print("Masukkan jam: ")
	fmt.Scan(&jam)
	switch {
	case jam == 0:
		waktu = "12 AM"
	case jam == 12:
		waktu = "12 PM"
	case jam > 0 && jam < 12:
		waktu = fmt.Sprintf("%d AM", jam)
	case jam > 12 && jam < 24:
		waktu = fmt.Sprintf("%d PM", jam-12)
	default:
		waktu = "Perintah tidak dapat berjalan"
	}
	fmt.Println("Konversi waktu:", waktu)
}

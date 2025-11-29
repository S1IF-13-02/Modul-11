package main

import "fmt"

func main() {
	var jam int
	fmt.Print("Masukkan Jam: ")
	fmt.Scanln(&jam)

	switch {
	case jam == 0 || jam == 24:
		fmt.Println("Jam setelah di konversi:", 12, "AM")
	case jam == 12:
		fmt.Println("Jam setelah di konversi:", 12, "PM")
	case jam >= 1 && jam <= 11:
		fmt.Println("Jam setelah di konversi:", jam, "AM")
	case jam >= 13 && jam <= 24:
		fmt.Println("Jam setelah di konversi:", jam-12, "PM")
	default:
		fmt.Println("Error: Input Jam dari 0-24")
	}
}
package main

import "fmt"

func main() {
	var pH float64

	fmt.Println("Masukan Kadar pH Air: ")
	fmt.Scan(&pH)

	switch {
	case pH >= 6.5 && pH <= 8.6:
		fmt.Println("Air layak diminum")
	case pH < 6.5 && pH > 8.6:
		fmt.Println("Air tidak layak minum")
	case pH < 0 || pH > 14:
		fmt.Println("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14.")
	default:
		fmt.Println("Input tidak valid")
	}

}

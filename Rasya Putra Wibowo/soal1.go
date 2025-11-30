package main

import (
	"fmt"
)

func main() {
	var pH float64
	fmt.Print("Masukkan nilai pH: ")
	fmt.Scan(&pH)

	switch {
	case pH < 0 || pH > 14:
		fmt.Println("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14.")
	case pH >= 6.5 && pH <= 8.6:
		fmt.Println("Air layak minum")
	default:
		fmt.Println("Air tidak layak minum")
	}
}


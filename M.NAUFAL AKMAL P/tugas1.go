package main

import "fmt"

func main() {
	var x float64
	fmt.Print("Masukan Kadar Air (pH): ")
	fmt.Scan(&x)

	switch {
	case x < 0 || x > 14:
		fmt.Print("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14.")
	case x < 6.5 || x > 8.6:
		fmt.Print("Air tidak layak minum")
	default:
		fmt.Print("Air layak minum")
	}
}
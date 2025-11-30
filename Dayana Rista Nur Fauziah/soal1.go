package main

import "fmt"

func main() {
	var ph float64

	fmt.Print("Masukkan nilai pH: ")
	fmt.Scan(&ph)

	if ph < 0 || ph > 14 {
		fmt.Println("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14.")
	} else if ph >= 6.5 && ph <= 8.6 {
		fmt.Println("Air layak minum")
	} else {
		fmt.Println("Air tidak layak minum")
	}
}

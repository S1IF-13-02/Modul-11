package main

import "fmt"

func main() {
	var kadar float64
	fmt.Print("Masukkan kadar pH: ")
	fmt.Scan(&kadar)

	switch {
	case kadar < 0 || kadar > 14:
		fmt.Println("Input tidak valid, rentang pH 0 - 14")
	case kadar >= 6.5 && kadar <= 8.6:
		fmt.Println("Air Layak Minum")
	default:
		fmt.Println("Air Tidak Layak Minum")
	}
}

package main

import "fmt"

func main() {
	var pH float64
	fmt.Print("Masukkan pH air: ")
	fmt.Scan(&pH)
	switch {
	case pH >= 6.5 && pH <= 8.6:
		fmt.Println("Air Layak Minum")
	case pH < 6.5 && pH > 8.6:
		fmt.Println("Air Tidak Layak Minum")
	case pH > 14:
		fmt.Println("Input tidak valid, rentang pH 0 - 14")
	}
}

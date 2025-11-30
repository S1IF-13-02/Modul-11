package main

import "fmt"

func main() {
	var nilaiPH float64

	fmt.Print("input nilai pH: ")
	fmt.Scan(&nilaiPH)

	switch {
	case nilaiPH >= 6.5 && nilaiPH <= 8.6:
		fmt.Println("Kondisi: Layak konsumsi")
	case nilaiPH < 6.5 || nilaiPH > 8.6 || nilaiPH <= 14:
		fmt.Println("Kondisi: Tidak layak konsumsi")
	case nilaiPH > 14:
		fmt.Println("Data tidak valid")
	}
}

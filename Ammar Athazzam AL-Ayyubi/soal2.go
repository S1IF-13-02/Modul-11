package main

import "fmt"

func main() {
	var jenis string
	var durasi int
	fmt.Print("(motor/mobil/truk) dan durasi: ")
	fmt.Scan(&jenis, &durasi)

	switch jenis {
	case "motor":
		biaya := 2000
		fmt.Printf("Rp.%d", biaya*durasi)
	case "mobil":
		biaya := 5000
		fmt.Printf("Rp.%d", biaya*durasi)
	case "truk":
		biaya := 8000
		fmt.Printf("Rp.%d", biaya*durasi)
	}
}

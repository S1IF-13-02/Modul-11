package main

import "fmt"

func main() {
	var waktu int

	fmt.Println("Masukan Waktu: ")
	fmt.Scan(&waktu)

	switch {
	case waktu == 0:
		fmt.Print("12 AM")
	case waktu == 12:
		fmt.Print("12 PM")
	case waktu > 0 && waktu < 12:
		fmt.Println(waktu, "AM")
	case waktu > 12 && waktu < 24:
		fmt.Println(waktu-12, "PM")
	default:
		fmt.Println("Waktu tidak valid")
	}
}

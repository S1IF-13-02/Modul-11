package main

import "fmt"

func main() {
	var b int

	fmt.Print("Masukan: ")
	fmt.Scan(&b)

	switch {
	case b == 0:
		fmt.Println("12 AM")
	case b == 12:
		fmt.Println("12 PM")
	case b > 0 && b < 12:
		fmt.Printf("%d AM\n", b)
	case b > 12 && b < 24:
		fmt.Printf("%d PM\n", b-12)
	default:
		fmt.Println("Angka yang anda masukkan tidak valid!")
	}
}

package main

import "fmt"

func main() {
	var t string
	fmt.Print("masukan nama tanaman: ")
	fmt.Scan(&t)

	switch t {
	case "nepenthes":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Tanaman asli Indonseia")
	case "venus":
		fmt.Println("termasuk Tanaman Karnivora")
		fmt.Println("Bukan asli Indonesia")
	default:
		fmt.Println("Tidak termasuk tanaman karnivora")
	}
}

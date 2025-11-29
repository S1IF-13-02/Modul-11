package main

import "fmt"

func main() {
	var tanaman string
	fmt.Print("Masukkan Nama Tanaman: ")
	fmt.Scan(&tanaman)
	switch tanaman {
	case "Nepenthes", "Utricularia_aurea":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Asli Indonesia")
	case "Venus", "Sarracenia", "Butterwort", "Drosera":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Bukan Asli Indonesia")
	default:
		fmt.Println("Tidak Termasuk Tanaman Karnivora")
	}
}

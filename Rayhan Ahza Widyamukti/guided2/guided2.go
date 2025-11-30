package main

import "fmt"

func main() {
	var tanaman string

	fmt.Println("Masukan Jenis Tanaman Karnivora: ")
	fmt.Scan(&tanaman)

	switch tanaman {
	case "nepenthes", "drosera", "utricularia", "aldrovanda":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Asli Indonesia")

	case "venus", "sarracenia", "darlingtonia":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Bukan Asli Indonesia")

	default:
		fmt.Println("Tidak Termasuk Tanaman Karnivora")
	}
}

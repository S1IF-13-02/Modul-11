package main

import (
	"fmt"
)

func main() {
	var tanaman string
	fmt.Print("Masukkan nama tanaman: ")
	fmt.Scan(&tanaman)


	switch tanaman {
	case "nepenthes", "drosera":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Asli Indonesia")
	case "venus", "pinguicula", "sarracenia", "darlingtonia":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Bukan Asli Indonesia")
	default:
		fmt.Println("Tidak termasuk Tanaman Karnivora")
	}
}

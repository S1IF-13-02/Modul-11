package main

import (
	"fmt"
	"strings"
)

func main() {
	var namaTanaman string

	fmt.Print("Masukkan nama tanaman: ")
	fmt.Scanln(&namaTanaman)

	namaTanaman = strings.ToLower(namaTanaman)

	switch namaTanaman {
	case "nepenthes":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Asal Indonesia")
	case "venus":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Asal Indonesia")
	case "karedok":
		fmt.Println("Tidak Termasuk Tanaman Karnivora")
	default:
		fmt.Println("Tanaman tidak dikenali")
	}
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var namaTanaman string
	fmt.Print("Masukkan nama tanaman: ")
	fmt.Scan(&namaTanaman)

	switch strings.ToLower(namaTanaman) {
	case "napenthes":
		fmt.Println("Termasuk Tanaman karnivora.")
		fmt.Println("Asli Indonesia.")
	case "venus":
		fmt.Println("Termasuk Tanaman karnivora.")
		fmt.Println("Bukan Asli Indonesia.")
	case "karedok":
		fmt.Println("Bukan Termasuk Tanaman karnivora.")

	}
}

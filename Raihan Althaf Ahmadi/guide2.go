package main

import "fmt"

func main() {
	var tanaman string
	fmt.Print("Nama Tanaman : ")
	fmt.Scan(&tanaman)

	jenis := ""
	asal := ""

	switch tanaman {
	case "kantong semar", "nepentes":
		jenis = "Tanaman Karnivora"
	case "venus", "sundew", "butterwort", "bladderwort":
		jenis = "Tanaman Karnivora"
	default:
		jenis = "Bukan Tanaman Karnivora"
	}

	switch tanaman {
	case "kantong semar", "nepentes":
		asal = "Asli Indonesia"
	default:
		asal = "Bukan Asli Indonesia"
	}

	fmt.Println(jenis, asal)
}

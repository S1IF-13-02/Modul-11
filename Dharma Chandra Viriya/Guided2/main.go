package main

import (
	"fmt"
	"strings"
)

func main() {
	var tumbuhanInput string

	fmt.Print("Masukan: ")
	fmt.Scan(&tumbuhanInput)

	tumbuhan := strings.ToLower(tumbuhanInput)

	switch tumbuhan {
	case "nepenthes", "drosera", "utricularia":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Asli Indonesia")
	case "venus", "pinguicula", "aldrovanda", "darlingtonia":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Bukan Asli Indonesia")
	default:
		fmt.Println("Tidak Termasuk Tanaman Karnivora")
	}
}

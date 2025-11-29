package main

import "fmt"

func main() {
	var tanaman string
	fmt.Scan(&tanaman)

	switch tanaman {
	case "nepenthes" :
		fmt.Println("Termasuk Tanaman Karnivora Asli Indonesia")
	case "venus", "Sarracenia", "Byblis", "Drosera" :
		fmt.Println("Termasuk Tanaman Karnivora Bukan Asli Indonesia")
	default:
		fmt.Println("Tidak Termasuk Tanaman Karnivora")
	}

}
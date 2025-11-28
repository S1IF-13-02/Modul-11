package main

import "fmt"

func main() {
	var tanaman string

	fmt.Scan(&tanaman)

	switch tanaman {
	case "nepenthes":
		fmt.Println("termasuk tanaman Karnivora Asli Indonesia")
		fmt.Println("asli indonesia")
	case "venus":
		fmt.Println("termasuk tanaman Karnivora Asli Indonesia")
		fmt.Println("bukan asli indonesia")
	default:
		fmt.Println("Tidak termasuk tanaman Karnivora")
	}

}

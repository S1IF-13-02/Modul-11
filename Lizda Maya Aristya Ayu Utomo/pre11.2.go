package main

import "fmt"

func main() {
	var tanaman string

	fmt.Scan(&tanaman)

	switch tanaman {
	case "nepenthes":
		fmt.Println("termasuk tanaman Karnivora Asli Indonesia")
	case "venus":
		fmt.Println("termasuk tanaman karnivora bukan asli indonrsia")
	default:
		fmt.Println("Tidak termasuk tanaman Karnivora")
	}

}

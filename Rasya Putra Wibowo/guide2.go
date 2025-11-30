package main

import "fmt"

func main() {
	var tanaman string

	fmt.Scan(&tanaman)

	switch tanaman {
	case "nepenthes", "kantong semar", "venus", "sarracenia":
		fmt.Println("Termasuk Tanaman Karnivora")

	switch tanaman {
	case "nepenthes", "kantong semar":
			fmt.Println("Asli Indonesia")
	case "venus", "sarracenia":
			fmt.Println("Bukan asli Indonesia")
		}
	default:
		fmt.Println("Tidak termasuk tanaman karnivora")
	}
}
package main

import(
	"fmt"
)

func main() {
	var tanaman string
	fmt.Print("Masukan Tanaman : ")
	fmt.Scan(&tanaman)

	switch tanaman{
	case "nepenthes", "drosera":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("asli indonesia")

	case "venus", "sarracenia":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Bukan asli indonesia")

	default :
		fmt.Println("Bukan Tanaman karnivora")
	}
}
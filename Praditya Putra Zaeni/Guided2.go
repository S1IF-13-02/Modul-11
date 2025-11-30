package main

import(
	"fmt"
)

func main() {
	var tanaman string
	fmt.Print("jenis Tanaman : ")
	fmt.Scan(&tanaman)

	switch tanaman{
	case "nepenthes", "drosea":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("asli indonesia")

	case "venus", "sarracenia", "butterwort":
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Bukan asli indonesia")

	default :
		fmt.Println("Bukan Tanaman karnivora")
	}
}
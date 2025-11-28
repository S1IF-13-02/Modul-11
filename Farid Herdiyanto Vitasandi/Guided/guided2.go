package main

import "fmt"

func main() {


	var tanaman string

	fmt.Print("Masukkan nama tanaman: ")
	fmt.Scan(&tanaman)

	switch tanaman{
	case "nepenthes", "drosera":
		fmt.Println("Termasuk tanaman karnivora")
		fmt.Println("Asli Indonesia")
	case "venus", "sarracenia":
		fmt.Println("Termasuk tanaman karnivora")
		fmt.Println("Bukan Asli Indonesia")
	default:
		fmt.Println("Tidak termasuk tanaman karnivora")
	}
}


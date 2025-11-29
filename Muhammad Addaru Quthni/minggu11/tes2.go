package main
import "fmt"
func main() {
	var tanaman string
	fmt.Print("Masukkan nama tanaman: ")
	fmt.Scanln(&tanaman)
	switch tanaman {
	case "nepenthes":
		fmt.Println("Termasuk Tanaman Karnivora.")
		fmt.Println("asli indonesia")
	case "venus":
		fmt.Println("Termasuk Tanaman Karnivora.")
		fmt.Println("bukan asli indonesia")
	default:
		fmt.Println("tidak termasuk tanaman karnivora.")
	}
}
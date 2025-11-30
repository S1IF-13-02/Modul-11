package main

import "fmt"
func main() {
	var kendaraan string
	var lama_parkir int

	fmt.Scan(&kendaraan)
	fmt.Scan(&lama_parkir)

	switch kendaraan {
	case "mobil":
		if lama_parkir <= 2 {
			fmt.Println("Biaya parkir: Rp 15000")
		} else {
			fmt.Printf("Biaya parkir: Rp 20000\n")
		}
	case "motor":
		if lama_parkir <= 2 {
			fmt.Println("Biaya parkir: Rp 7000")
		} else {
			fmt.Printf("Biaya parkir: Rp9000\n")
		}
	case "truck":
		if lama_parkir <= 2 {
			fmt.Println("Biaya parkir: Rp 25000")
		} else {
			fmt.Printf("Biaya parkir: Rp35000\n")
		}
	default:
		fmt.Println("Jenis kendaraan tidak dikenali")
	}
}
package main

import "fmt"

func main() {
	var angka int
	fmt.Scan(&angka)

	switch {
	case angka%10 == 0:
		fmt.Printf("Kategori: Bilangan Kelipatan 10\nHasil pembagian antara %d / 10 = %d\n", angka, angka/10)
	case angka%5 == 0 && angka > 5:
		fmt.Printf("Kategori: Bilangan Kelipatan 5\nHasil kuadrat dari %d ^2 = %d\n", angka, angka*angka)
	case angka%2 == 0:
		fmt.Printf("Kategori: Bilangan Genap\nHasil perkalian dengan bilangan berikutnya %d * %d = %d\n", angka, angka+1, angka*(angka+1))
	default:
		fmt.Printf("Kategori: Bilangan Ganjil\nHasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", angka, angka+1, angka+(angka+1))
	}
}

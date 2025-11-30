package main

import (
	"fmt"
)

func main() {
	var input int
	var hasil int

	fmt.Print("Masukkan satu bilangan bulat: ")
	fmt.Scan(&input)
	switch {
	case input%10 == 0:
		hasil = input / 10
		fmt.Printf("Kategori: Bilangan Kelipatan 10\nDilanjutkan dengan \"Hasil pembagian antara %d / 10 = %d\" untuk Bilangan Kelipatan 10.\n", input, hasil)
	case input%5 == 0:
		hasil = input * input
		fmt.Printf("Kategori: Bilangan Kelipatan 5\nDilanjutkan dengan \"Hasil kuadrat dari %d ^ 2 = %d\" untuk Bilangan Kelipatan 5.\n", input, hasil)
	case input%2 == 0:
		hasil = input * (input + 1)
		fmt.Printf("Kategori: Bilangan Genap\nDilanjutkan dengan \"Hasil perkalian dengan bilangan berikutnya %d * %d = %d\" untuk Bilangan Genap.\n", input, input+1, hasil)
	default:
		hasil = input + (input + 1)
		fmt.Printf("Kategori: Bilangan Ganjil\nDilanjutkan dengan \"Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\" untuk Bilangan Ganjil.\n", input, input+1, hasil)
	}
}

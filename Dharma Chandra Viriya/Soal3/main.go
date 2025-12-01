package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Print("Masukan: ")
	fmt.Scan(&n)

	switch {
	case n%10 == 0:
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		fmt.Printf("Hasil kuadrat dari %d / 10 = %d\n", n, n/10)
	case n%5 == 0 && n != 5:
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		fmt.Printf("Hasil kuadrat dari %d ^2 = %d\n", n, int(math.Pow(float64(n), 2)))
	case n%2 == 0:
		fmt.Println("Kategori: Bilangan Genap")
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d + %d = %d\n", n, n+1, n*(n+1))
	case n%2 != 0:
		fmt.Println("Kategori: Bilangan Ganjil")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", n, n+1, n+(n+1))
	}
}

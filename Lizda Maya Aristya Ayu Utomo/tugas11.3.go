package main

import "fmt"

func main() {
	var b int
	fmt.Scan(&b)

	switch {
	case b%10 == 0:
		kelipatan10 := b / 10
		fmt.Println("kategori bilangan kelipatan 10")
		fmt.Printf("hasil pembagian antara %d / 10 = %d", b, kelipatan10)

	case b&5 == 0 && b != 5:
		kelipatan5 := b * b
		fmt.Println("kategori: bilangan kelipatan 5")
		fmt.Printf("hasil kuadrat dari %d^2 = %d", b, kelipatan5)

	case b%2 == 0:
		bilangangenap := b * (b + 1)
		fmt.Println("kategori: bilangan genap")
		fmt.Printf("hasil perkalian dengan bilangan berikutnya %d * %d = %d", b, b+1, bilangangenap)

	default:
		bilanganganjil := b + (b + 1)
		fmt.Println("kategori: bilangan ganjil")
		fmt.Printf("hasil penjumlahan dengan bilangan berikutnya %d + %d = %d", b, b+1, bilanganganjil)

	}
}
package main

import "fmt"

func main() {
	var n int

	fmt.Print("masukan bilangan : ")
	fmt.Scan(&n)

	switch {
	case n%10 == 0:
		hasil := n / 10
		fmt.Println("Kategori : Bilangan kelipatan 10")
		fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", n, hasil)
	case n%5 == 0 && n != 5:
		hasil := n * n
		fmt.Println("Kategori : Bilangan kelipatan 5")
		fmt.Printf("Hasil kuadrat dari %d ^ 2 = %d\n", n, hasil)
	case n%2 != 0:
		berikutnya := n + 1
		hasil := n + berikutnya
		fmt.Println("Kategori : Bilangan Ganjil")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", n, berikutnya, hasil)
	default:
		berikutnya := n + 1
		hasil := n * berikutnya
		fmt.Println("Kategori : Bilangan Genap")
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n", n, berikutnya, hasil)
	}
}
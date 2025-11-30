package main

import "fmt"

func main() {
	var x int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&x)

	if x%10 == 0 {
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		fmt.Println("Hasil pembagian antara", x, "/ 10 =", x/10)

	} else if x == 5 {
		fmt.Println("Kategori: Bilangan Ganjil")
		fmt.Println("Hasil penjumlahan dengan bilangan berikutnya", x, "+", (x + 1), "=", x+(x+1))

	} else if x%5 == 0 {
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		fmt.Println("Hasil kuadrat dari", x, "^ 2 =", x*x)

	} else if x%2 == 0 {
		fmt.Println("Kategori: Bilangan Genap")
		fmt.Println("Hasil perkalian dengan bilangan berikutnya", x, "*", (x + 1), "=", x*(x+1))

	} else {
		fmt.Println("Kategori: Bilangan Ganjil")
		fmt.Println("Hasil penjumlahan dengan bilangan berikutnya", x, "+", (x + 1), "=", x+(x+1))
	}
}

package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	switch {
	case bilangan % 10 == 0 :
		kelipatan10 := bilangan / 10
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		fmt.Printf("Hasil pembagian antara %d / 10 = %d",bilangan,kelipatan10)
	case bilangan % 5 == 0 && bilangan != 5 :
		kelipatan5 := bilangan * bilangan
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		fmt.Printf("Hasil kuadrat dari %d^2 = %d",bilangan,kelipatan5)
	case bilangan % 2 == 0 :
		bilangangenap := bilangan * (bilangan + 1)
		fmt.Println("Kategori: Bilangan Genap")
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d",bilangan,bilangan+1,bilangangenap)
	default :
		bilanganganjil := bilangan + (bilangan + 1)
		fmt.Println("Kategori: Bilagan Ganjil")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d", bilangan,bilangan+1,bilanganganjil)
	}
}
package main

import "fmt"

func main() {
	
	var bilangan int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)

	switch{
	case bilangan == 0:
		fmt.Println("Input tidak boleh nol")
	case bilangan%10 == 0:
		hasil := bilangan / 10
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		fmt.Printf("Hasil pembagian antara %d dengan 10 adalah %d", bilangan, hasil)
	case bilangan%5 == 0:
		hasil := bilangan * bilangan
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		fmt.Printf("Hasil kuadrat dari %d adalah %d", bilangan, hasil)
	case bilangan%2 == 0:
		hasil := bilangan * (bilangan + 1)
		fmt.Println("Kategori: Bilangan Genap")
		fmt.Printf("Hasil perkalian antara %d dengan bilangan berikutnya %d adalah %d", bilangan, bilangan+1, hasil)
	case bilangan%2 != 0:
		hasil := bilangan + (bilangan + 1)
		fmt.Println("Kategori: Bilangan Ganjil")
		fmt.Printf("Hasil penjumlahan antara %d dengan bilangan berikutnya %d adalah %d", bilangan, bilangan+1, hasil)
	}
}

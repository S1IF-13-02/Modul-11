package main

import "fmt"

func main () {
	var bilangan, total int
	fmt.Print("Masukkan bilangan = ")
	fmt.Scan(&bilangan)

switch {
	case bilangan%10 == 0:
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		total = bilangan / 10
		fmt.Println("Hasil pembagian antara", bilangan, "/ 10 =", total)

	case bilangan%2 == 0:
		fmt.Println("Kategori: Bilangan Genap")
		total = bilangan * (bilangan + 1)
		fmt.Println("Hasil perkalian antara", bilangan, "*", bilangan+1, "=", total)

	case bilangan%5 == 0 && bilangan != 5:
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		total = bilangan * bilangan
		fmt.Println("Hasil kuadrat dari", bilangan, "^2 =", total)

	case bilangan%2 != 0:
		fmt.Println("Kategori: Bilangan Ganjil")
		total = bilangan + (bilangan + 1)
		fmt.Println("Hasil penjumlahan antara", bilangan, "+", bilangan+1, "=", total)

	default:
		fmt.Println("Tidak ada kategori yang sesuai")
	}	
}
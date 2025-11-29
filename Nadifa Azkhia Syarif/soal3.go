package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	next := n + 1
	var hasil int

	switch {
	case n%10 == 0:
		fmt.Println("Kelipatan 10")
		hasil = n / 10
		fmt.Println("Hasil:", hasil)

	case n%5 == 0:
		fmt.Println("Kelipatan 5")
		hasil = n * n
		fmt.Println("Hasil:", hasil)

	case n%2 == 0:
		fmt.Println("Genap")
		hasil = n * next
		fmt.Println("Hasil:", hasil)

	default:
		fmt.Println("Ganjil")
		hasil = n + next
		fmt.Println("Hasil:", hasil)
	}
}

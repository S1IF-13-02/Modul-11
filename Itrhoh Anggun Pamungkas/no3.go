package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&n)

    switch {
    case n%10 == 0: 
        fmt.Println("Kategori: Bilangan Kelipatan 10")
        hasil := n / 10
        fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", n, hasil)

    case n%5 == 0: 
        fmt.Println("Kategori: Bilangan Kelipatan 5")
        hasil := n * n
        fmt.Printf("Hasil kuadrat dari %d ^ 2 = %d\n", n, hasil)

    case n%2 == 1: 
        fmt.Println("Kategori: Bilangan Ganjil")
        hasil := n + (n + 1)
        fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n",
            n, n+1, hasil)

    case n%2 == 0: 
        fmt.Println("Kategori: Bilangan Genap")
        hasil := n * (n + 1)
        fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n",
            n, n+1, hasil)

    default:
        fmt.Println("Input tidak valid")
    }
}
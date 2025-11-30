package main

import "fmt"

func main() {
    var input int
    fmt.Scan(&input)

    switch {
    case input%10 == 0:
        hasil := input / 10
        fmt.Println("Kategori: Bilangan Kelipatan 10")
        fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", input, hasil)

    case input%5 == 0:
        hasil := input * input
        fmt.Println("Kategori: Bilangan Kelipatan 5")
        fmt.Printf("Hasil kuadrat dari %d ^2 = %d\n", input, hasil)

    case input%2 == 0:
        hasil := input * (input + 1)
        fmt.Println("Kategori: Bilangan Genap")
        fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n",
            input, input+1, hasil)

    default:
        hasil := input + (input + 1)
        fmt.Println("Kategori: Bilangan Ganjil")
        fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n",
            input, input+1, hasil)
    }
}

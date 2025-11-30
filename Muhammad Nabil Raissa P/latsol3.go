package main

import "fmt"

func main() {
    var inputAngka int

    fmt.Print("Masukkan bilangan bulat: ")
    fmt.Scan(&inputAngka)

    angkaSelanjutnya := inputAngka + 1
    totalTambah := inputAngka + angkaSelanjutnya
    totalKali := inputAngka * angkaSelanjutnya
    nilaiKuadrat := inputAngka * inputAngka
    hasilBagi10 := inputAngka / 10

    switch {
    case inputAngka%10 == 0:
        fmt.Println("Kategori: Bilangan Kelipatan 10")
        fmt.Printf("Hasil pembagian %d / 10 = %d\n", inputAngka, hasilBagi10)

    case inputAngka%5 == 0 && inputAngka != 5:
        fmt.Println("Kategori: Bilangan Kelipatan 5")
        fmt.Printf("Hasil kuadrat dari %d = %d\n", inputAngka, nilaiKuadrat)

    case inputAngka%2 == 0:
        fmt.Println("Kategori: Bilangan Genap")
        fmt.Printf("Hasil perkalian %d * %d = %d\n", inputAngka, angkaSelanjutnya, totalKali)

    default:
        fmt.Println("Kategori: Bilangan Ganjil")
        fmt.Printf("Hasil penjumlahan %d + %d = %d\n", inputAngka, angkaSelanjutnya, totalTambah)
    }
}

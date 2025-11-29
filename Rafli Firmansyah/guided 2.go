package main

import "fmt"

func main () {
    var nama string         

    fmt.Print("Masukkan nama tanaman: ")
    fmt.Scan(&nama)          

    switch nama {
    case "nepenthes":        
        fmt.Println("Termasuk Tanaman Karnivora")
        fmt.Println("Asli Indonesia")

    case "venus":          
        fmt.Println("Termasuk Tanaman Karnivora")
        fmt.Println("Bukan Asli Indonesia")

    default:                
        fmt.Println("Tidak termasuk Tanaman Karnivora")
    }
}
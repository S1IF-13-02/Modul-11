package main

import "fmt"

func main() {
    var jenis string
    var durasi int

    fmt.Print("Jenis kendaraan (Motor/Mobil/Truk): ")
    fmt.Scan(&jenis)
    fmt.Print("Durasi parkir (jam): ")
    fmt.Scan(&durasi)

    switch jenis {
    case "Motor":
        if durasi <= 2 {
            fmt.Println("Rp 7000")
        } else {
            fmt.Println("Rp 9000")
        }
    case "Mobil":
        if durasi <= 2 {
            fmt.Println("Rp 15000")
        } else {
            fmt.Println("Rp 20000")
        }
    case "Truk":
        if durasi <= 2 {
            fmt.Println("Rp 25000")
        } else {
            fmt.Println("Rp 35000")
        }
    default:
        fmt.Println("Jenis kendaraan atau durasi waktu tidak valid")
    }
}

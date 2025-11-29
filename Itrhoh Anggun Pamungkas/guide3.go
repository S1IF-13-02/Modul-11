package main

import "fmt"

func main() {
    var jenis string
    var durasi, tarif int

    fmt.Print("jenis kendaraan (motor/mobil/truk): ")
    fmt.Scanln(&jenis)

    fmt.Print("durasi parkir (dalam jam): ")
    fmt.Scanln(&durasi)

    switch jenis {
    case "motor":
        if durasi >= 1 && durasi <= 2 {
            tarif = 7000
        } else if durasi > 2 {
            tarif = 9000
        }

    case "mobil":
        if durasi >= 1 && durasi <= 2 {
            tarif = 15000
        } else if durasi > 2 {
            tarif = 20000
        }

    case "truk":
        if durasi >= 1 && durasi <= 2 {
            tarif = 25000
        } else if durasi > 2 {
            tarif = 30000
        }

    default:
        fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
		return
    }

    fmt.Println("Tarif Parkir: Rp", tarif)
}

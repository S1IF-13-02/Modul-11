package main
import "fmt"

func main() {
    var jenis string
    var jam, tarif int

    fmt.Scan(&jenis, &jam)

    if jam < 1 {
        jam = 1
    }

    switch jenis {
    case "motor":
        tarif = 2000 * jam
    case "mobil":
        tarif = 5000 * jam
    case "truk":
        tarif = 8000 * jam
    default:
        fmt.Println("Jenis kendaraan tidak valid")
        return
    }

    fmt.Println("Rp", tarif)
}

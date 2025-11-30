package main

import "fmt"

func main() {
    var jam int
    fmt.Print("Masukan jam: ")
    fmt.Scan(&jam)
    
    switch {
    case jam == 0:
        fmt.Print("12 AM")
    case jam == 12:
        fmt.Print("12 PM")
    case jam > 0 && jam < 12:
        fmt.Println(jam, "AM")
    case jam > 12 && jam < 24:
        fmt.Println(jam-12, "PM")
    default:
        fmt.Println("Waktu tidak valid")
}
}

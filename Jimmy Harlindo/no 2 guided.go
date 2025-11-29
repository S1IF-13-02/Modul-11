package main

import (
    "fmt"
    "strings"
)

func main() {
    var tanaman string

    fmt.Print("masukan nama tanaman: ")
    fmt.Scanln(&tanaman)

    t := strings.ToLower(tanaman)

    switch t {
    case "nepenthes":
        fmt.Println("termasuk tanaman karnivora")
        fmt.Println("asli indonesia")

    case "venus":
        fmt.Println("termasuk tanaman karnivora")
        fmt.Println("bukan asli indonesia")

    default:
        fmt.Println("tidak termasuk tanaman karnivora")
    }
}

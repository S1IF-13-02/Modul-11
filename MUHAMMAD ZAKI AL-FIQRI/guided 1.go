package main

import "fmt"

func main() {

	var n int
	fmt.Print("Masukan jam : ")
	fmt.Scan(&n)

	switch {
	case n > 0 && n < 12:
		fmt.Println(n, " AM")
	case n > 12 && n <= 24:
		fmt.Println(n-12, " PM")
	case n == 0:
		fmt.Println(12, " AM")
	case n == 12:
		fmt.Println(12, " PM")
	default:
		fmt.Println("inputan salah")
	}
}

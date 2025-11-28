package main

import "fmt"

func main() {
	var a int

	fmt.Print("masukan jam: ")
	fmt.Scan(&a)

	switch {
	case a == 0:
		fmt.Println("12 AM")
	case a == 12:
		fmt.Println("12 PM")
	case a > 0 && a < 12:
		fmt.Println(a, "AM")
	case a > 12 && a < 24:
		fmt.Println(a-12, "PM")
	default:
		fmt.Println("tidak valid")
	}
}

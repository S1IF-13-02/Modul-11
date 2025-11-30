package main

import "fmt"

func main() {
	var jam24 int
	fmt.Print("Masukkan jam 0-23: ")
	fmt.Scan(&jam24)

	switch {
	case jam24 == 0:
		fmt.Println(12, "AM")

	case jam24 == 12:
		fmt.Println(12, "PM")

	case jam24 > 12:
		jam12 := jam24 - 12
		fmt.Println(jam12, "PM")

	default:
		jam12 := jam24
		fmt.Println(jam12, "AM")
	}
}

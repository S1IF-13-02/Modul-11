package main

import (
	"fmt"
)

func main() {
	var jam24 int
	fmt.Scan(&jam24)

	var jam12 int
	var kondisi string

	switch {
	case jam24 == 0:
		jam12 = 12
		kondisi = "AM"
	case jam24 == 12:
		jam12 = 12
		kondisi = "PM"
	case jam24 >= 1 && jam24 <= 11:
		jam12 = jam24
		kondisi = "AM"
	case jam24 >= 13 && jam24 <= 23:
		jam12 = jam24 - 12
		kondisi = "PM"
	default:
		fmt.Println("Input harus Jam 0-23")
	}

	fmt.Printf("%d %s\n", jam12, kondisi)
}
package main

import (
	"fmt"
)

func main() {
	var jam24,jam12 int	
	var ampm string

	fmt.Print("Masukkan jam : ")
	fmt.Scan(&jam24)

	switch {
	case jam24 == 0:
		jam12 = 12
		ampm = "AM"
	case jam24 == 12:
		jam12 = 12
		ampm = "PM"
	case jam24 >= 1 && jam24 <= 11:
		jam12 = jam24
		ampm = "AM"
	case jam24 >= 13 && jam24 <= 23:	
		jam12 = jam24 - 12
		ampm = "PM"
	default:
		fmt.Println("invalid, Input harus Jam 0-23")
		return
	}

	fmt.Printf("%d %s\n", jam12, ampm)
}

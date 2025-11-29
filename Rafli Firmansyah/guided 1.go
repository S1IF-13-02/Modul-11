package main

import "fmt"

func main() {
	var jam24, jam12 int
	var ket string
	fmt.Scan(&jam24)

	switch {
	case jam24 == 0:
		jam12 = 12
		ket = "AM"

	case jam24 == 12:
		jam12 = 12
		ket = "PM"

	case jam24 >= 1 && jam24 <= 11:
		jam12 = jam24
		ket = "AM"

	case jam24 >= 13 && jam24 <= 23:
		jam12 = jam24 - 12
		ket = "PM"

	default:
		fmt.Println("Jam eror")
		return
	}

	fmt.Printf("%d %s\n", jam12, ket)
}

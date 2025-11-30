package main

import "fmt"

func main() {
	var hour24 int

	fmt.Print("Masukkan jam (0-23): ")
	fmt.Scan(&hour24)

	if hour24 < 0 || hour24 > 23 {
		fmt.Println("Error: Jam harus antara 0-23")
		return
	}

	var hour12 int
	var period string

	switch {
	case hour24 == 0:
		hour12 = 12
		period = "AM"
	case hour24 < 12:
		hour12 = hour24
		period = "AM"
	case hour24 == 12:
		hour12 = 12
		period = "PM"
	default:
		hour12 = hour24 - 12
		period = "PM"
	}

	fmt.Printf("\nFormat 24 jam: %02d:00\n", hour24)
	fmt.Printf("Format 12 jam: %02d:00 %s\n", hour12, period)
}

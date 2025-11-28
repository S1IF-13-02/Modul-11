package main

import "fmt"

func main() {

	var jam24, jam12 int
	var periode string

	fmt.Print("Masukkan jam dalam format 24 jam (0–23): ")
	fmt.Scan(&jam24)

	if jam24 < 0 || jam24 > 23 {
		fmt.Println("Input tidak valid! Masukkan angka 0–23.")
		return
	}

	switch {
	case jam24 == 0:
		jam12 = 12
		periode = "AM"
	case jam24 == 12:
		jam12 = 12
		periode = "PM"
	case jam24 > 12:
		jam12 = jam24 - 12
		periode = "PM"
	default:
		jam12 = jam24
		periode = "AM"
	}

	fmt.Printf("Hasil konversi: %d %s\n", jam12, periode)
}

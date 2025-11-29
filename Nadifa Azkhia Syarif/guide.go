package main

import "fmt"

func main() {
	var jam24 int

	fmt.Println("Masukkan jam (0-23): ")
	fmt.Scan(&jam24)

	switch {
	case jam24 == 0:
		fmt.Println("hasil konverensi: ")
		fmt.Println("12 AM")

	case jam24 == 12:
		fmt.Println("hasil konverensi: ")
		fmt.Println("12 PM")

	case jam24 > 0 && jam24 < 12:
		fmt.Println("hasil konverensi: ")
		fmt.Println(jam24, "AM")

	case jam24 > 12 && jam24 < 24:
		fmt.Println("hasil konverensi: ")
		fmt.Println(jam24-12, "PM")

	default:
		fmt.Println("Error: pilihan tidak valid.")

	}
}

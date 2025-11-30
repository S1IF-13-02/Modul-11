package main

import "fmt"

func main() {
	var jam int
	fmt.Scan(&jam)

	if jam < 0 || jam > 23 {
		fmt.Println("Jam tidak valid")
		return
	}

	switch {
		case jam == 0:
			fmt.Println("12 AM")
		case jam == 12:
			fmt.Println("12 PM")
		case jam > 12:
			fmt.Printf("%d PM\n", jam-12)
		default:
			fmt.Printf("%d AM\n", jam)
	}
}
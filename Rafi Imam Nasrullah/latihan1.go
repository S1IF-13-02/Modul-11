package main

import "fmt"

func main() {
	var jam int

	fmt.Scan(&jam)

	switch  {
	case jam < 0 || jam > 23 : 
		fmt.Println("Jam Tidak Valid")
	case jam == 0 :
		fmt.Println("12 AM")
	case jam == 12 :
		fmt.Println("12 PM")
	case jam < 12 :
		fmt.Println(jam, " AM")
	default : 
		fmt.Println(jam - 12, " PM")
	}
}

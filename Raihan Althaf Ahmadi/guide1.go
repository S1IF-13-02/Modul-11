package main

import "fmt"

func main() {
	var jam int
	fmt.Print("Masukan Jam (0-23): ")
	fmt.Scan(&jam)

	switch {
	case jam < 0 || jam > 23:
		fmt.Println("Jam tidak valid")
	case jam == 0:
		fmt.Println("12 AM") 
	case jam == 12:
		fmt.Println("12 PM") 
	case jam < 12:
		fmt.Printf("%d AM\n", jam)
	case jam > 12 && jam < 18  :
		fmt.Printf("%d PM\n", jam-12)
	case jam < 12 && jam < 6 :
		fmt.Printf("%d AM\n", jam-12)
	default:
		fmt.Printf("%d AM\n", jam-12)
	}
}
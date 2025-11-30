package main

import "fmt"

func main() {
	var ph float64

	fmt.Print("masukan kadar pH: ")
	fmt.Scan(&ph)

	switch {
	case ph >= 6.5 && ph <= 8.6:
		fmt.Println("Air layak minum")
	case ph < 6.5 || ph > 8.6:
		fmt.Println("Air tidak layak minum")
	case ph > 14:
		fmt.Println("Input tidak valid")
	}
}
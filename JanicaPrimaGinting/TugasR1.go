package main

import "fmt"

func main(){
	var ph float64
	fmt.Scan(&ph)
	switch {
	case ph >= 6.5 && ph <= 8.6 :
		fmt.Println("Air layak minum")
	case ph < 6.5 || ph > 8.6 && ph <= 14 :
		fmt.Println("Air tidak layak minum")
	default :
		fmt.Println("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14")
	}
}
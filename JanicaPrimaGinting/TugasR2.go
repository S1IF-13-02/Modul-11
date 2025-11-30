package main

import "fmt"

func main(){
	var kendaraan string
	var durasi int
	fmt.Scan(&kendaraan)
	fmt.Scan(&durasi)
	biaya := 0
	if durasi < 1 {
		durasi = 1
	}
	switch kendaraan {
	case "motor" :
		biaya = 2000
	case "mobil" :
		biaya = 5000
	case "truk" :
		biaya = 8000
	default :
		fmt.Println("Kendaraan Tidak Valid")
	
	}
	fmt.Println("Rp",biaya * durasi)
}
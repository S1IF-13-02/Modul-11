package main 

import "fmt"

func main() {
	
	var kadar_pH float64

	fmt.Print("Masukkan kadar pH air (0-14): ")
	fmt.Scan(&kadar_pH)

	if kadar_pH < 0 || kadar_pH > 14{
		fmt.Println("Nilai pH tidak valid. Nilai pH harus antara 0 - 14.")
		return
	}

	switch {
	case kadar_pH >= 6.5 && kadar_pH <= 8.6:
		fmt.Println("Air layak minum")
	default:
		fmt.Println("Air tidak layak minum")
	}
}

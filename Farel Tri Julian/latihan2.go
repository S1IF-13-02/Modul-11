package main

import "fmt"

func main() {
	var kendaraan string
	var jam int

	fmt.Scan(&kendaraan, &jam)

	switch kendaraan {
	case "motor":
		tarif := 2000 * jam
		fmt.Println("Rp", tarif)

	case "mobil":
		tarif := 5000 * jam
		fmt.Println("Rp", tarif)

	case "truk":
		tarif := 8000 * jam
		fmt.Println("Rp", tarif)

	}
}

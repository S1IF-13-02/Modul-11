package main

import "fmt"

func main() {
	var kendaraan string
	var waktu int
	fmt.Print("Masukkan jenis kendaraan: ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukkan lama waktunya: ")
	fmt.Scan(&waktu)
	switch kendaraan {
	case "motor":
		if waktu >= 1 && waktu <= 2 {
			fmt.Println("Tarif parkir: Rp. 7. 000")
		} else if waktu > 2 {
			fmt.Println("Tarif parkir: Rp. 9. 000")
		}
	case "mobil":
		if waktu >= 1 && waktu <= 2 {
			fmt.Println("Tarif parkir: Rp. 15. 000")
		} else if waktu > 2 {
			fmt.Println("Tarif parkir: Rp. 20. 000")
		}
	case "truk":
		if waktu >= 1 && waktu <= 2 {
			fmt.Println("Tarif parkir: Rp. 25. 000")
		} else if waktu > 2 {
			fmt.Println("Tarif parkir: Rp. 35. 000")
		}
	default:
		fmt.Println("Tidak dikenakan tarif")
	}
}

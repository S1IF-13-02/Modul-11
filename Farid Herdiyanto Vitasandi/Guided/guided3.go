package main

import "fmt"

func main() {

	var kendaraan string
	var lama_parkir int

	fmt.Print("Masukkan jenis kendaraan (Mobil/Motor/Truk): ")
	fmt.Scan(&kendaraan)

	fmt.Print("Masukkan durasi waktu parkir (jam): ")
	fmt.Scan(&lama_parkir)

	switch kendaraan{
	case "mobil":
		if lama_parkir <= 2{
		fmt.Println("Rp.15.000")
		} else {
		fmt.Println("Rp.20.000")
		}
	case "motor":
		if lama_parkir <= 2{
		fmt.Println("Rp.7.000")
		}else{
			fmt.Println("Rp.9.000")
		}
	case "truk":
		if lama_parkir <= 2{
		fmt.Println("Rp.25.000")
		}else{
			fmt.Println("Rp.35.000")
		}
	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
	}
	
}

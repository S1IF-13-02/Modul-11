package main 

import "fmt"

func main (){
	var kendaraan string 
	var durasi int
	var tarif int
	fmt.Scan(&kendaraan)
	fmt.Scan(&durasi)
	fmt.Println("Masukkan jenis kendaraan (Motor/Mobil/Truk):",kendaraan)
	fmt.Println("Masukkan durasi parkir (dalam jam):",durasi)
	switch {
	case kendaraan == "Motor" && durasi >= 1 && durasi <= 2 :
		tarif = 7000
	case kendaraan == "Motor" && durasi > 2 :
		tarif = 9000
	case kendaraan == "Mobil" && durasi >= 1 && durasi <= 2 :
		tarif = 15000
	case kendaraan == "Mobil" && durasi > 2 :
		tarif = 20000
	case kendaraan == "Truk" && durasi >= 1 && durasi <= 2 :
		tarif = 25000
	case kendaraan == "Truk" && durasi > 2 :
		tarif = 35000
	default :
		fmt.Println("Jenis Kendaraan atau Durasi Parkir tidak Valid")
	}
	fmt.Println("Tarif Parkir: Rp",tarif)
}
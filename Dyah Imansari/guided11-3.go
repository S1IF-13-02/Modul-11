package main
import "fmt"
func main() {
	var tarif int
	var durasi int
	var kendaraan string
	fmt.Print("Masukkan jenis kendaraan (Motor/Mobil/Truk): ")
	fmt.Scanln(&kendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scanln(&durasi)
	switch {
	case kendaraan == "Motor" && durasi >= 1 && durasi <= 2:
		tarif = 7000
	case kendaraan == "Motor" && durasi > 2:
		tarif = 9000
	case kendaraan == "Mobil" && durasi >= 1 && durasi <= 2:
		tarif = 15000
	case kendaraan == "Mobil" && durasi > 2:
		tarif = 20000
	case kendaraan == "Truk" && durasi >= 1 && durasi <= 2:
		tarif = 25000
	case kendaraan == "Truk" && durasi > 2:
		tarif = 35000
	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
	}
	fmt.Printf("Tarif Parkir: Rp %d\n", tarif)
}
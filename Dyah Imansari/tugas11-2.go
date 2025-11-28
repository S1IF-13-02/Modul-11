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
	if durasi < 1 {
		durasi = 1
	}
	switch {
	case kendaraan == "Motor":
		tarif = 2000*durasi
	case kendaraan == "Mobil":
		tarif = 5000*durasi
	case kendaraan == "Truk":
		tarif = 8000*durasi
	default:
		fmt.Println("Jenis kendaraan tidak valid")
	}
	fmt.Printf("Tarif Parkir untuk %d jam: Rp %d\n", durasi, tarif)
}
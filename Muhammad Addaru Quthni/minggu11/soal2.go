package main
import ("fmt"
	"strings")
func main() {
	var kendaraan string
	var durasi int
	fmt.Print("Masukkan jenis kendaraan (mobil/motor/truk): ")
	fmt.Scanln(&kendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scanln(&durasi)
	kendaraan = strings.ToLower(kendaraan)
	switch kendaraan {
	case "motor":
		tarif := 2000 * durasi
		fmt.Println("Biaya parkir motor: Rp.",tarif,"Jam")
	case "mobil":
		tarif := 5000 * durasi
		fmt.Println("Biaya parkir mobil: Rp.",tarif,"Jam")
	case "truk":
		tarif := 8000 * durasi
		fmt.Println("Biaya parkir truk: Rp.",tarif,"Jam")
	default:
		fmt.Println("Masukkan Jenis Kendaraan yang sesuai.")
	}
}
package main
import ( "fmt" 
	"strings"
)
func main() {
	var kendaraan string
	var durasi int
	fmt.Print("Masukkan jenis kendaraan (mobil/motor/truk): ")
	fmt.Scanln(&kendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scanln(&durasi)
	kendaraan = strings.ToLower(kendaraan)
	switch kendaraan {
	case "mobil":
		if durasi <= 2 {
			fmt.Println("Biaya parkir mobil: Rp 15.000")
		} else {
			fmt.Println("Biaya parkir mobil: Rp 20.000")
		}
	case "motor":
		if durasi <= 2 {
			fmt.Println("Biaya parkir motor: Rp 7.000")
		} else {
			fmt.Println("Biaya parkir motor: Rp 9.000")
		}
	case "truk":
		if durasi <= 2 {
			fmt.Println("Biaya parkir truk: Rp 25.000")
		} else {
			fmt.Println("Biaya parkir truk: Rp 35.000")	
		}
	default:
		fmt.Println("Masukkan Jenis Kendaraan yang sesuai.")
}
}
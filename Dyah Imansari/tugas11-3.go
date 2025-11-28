package main
import "fmt"
func main() {
	var x int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&x)
	switch {
	case x%10 == 0:
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", x, x/10)
	case x%5 == 0 && x!= 5:
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		fmt.Printf("Hasil kuadrat dari %d^2 = %d\n", x, x*x)
	case x%2 != 0:
		fmt.Println("Kategori: Bilangan Ganjil")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", x, x+1, x+(x+1))
	case x%2 == 0:
		fmt.Println("Kategori: Bilangan Genap")
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n", x, x+1, x*(x+1))
	}
}
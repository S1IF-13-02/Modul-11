package main
import "fmt"
func main() {
	var h int
	var jam int = 24
	var waktu string
	fmt.Scan(&jam)
	switch {
	case jam == 0:
		h = 12
		waktu = "AM"
	case jam >= 1 && jam <= 11:
		h = jam
		waktu = "AM"
	case jam == 12:
		h = 12
		waktu = "PM"
	case jam >= 13 && jam <= 23:
		h = jam - 12
		waktu = "PM"
	case jam == 24:
		h = 12
		waktu = "AM"
	}
	fmt.Println(h, waktu)
}
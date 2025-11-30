
package main

import "fmt"

func main() {
	var jam int
	fmt.Scan(&jam)
	var konversi int
	var periode string
	switch (jam) {
	case 0:
		konversi = 12;
		periode = "AM";
	case 12:
		konversi = 12;
		periode = "PM"
	default:
		if (jam < 24){
			konversi = jam - 12;
			periode = "PM"
		}else{
			konversi = jam
			periode = "AM"
		}
	}
	fmt.Println(konversi, periode)
}
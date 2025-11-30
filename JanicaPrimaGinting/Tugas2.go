package main 

import "fmt"

func main(){
	var tnmn string
	fmt.Scan(&tnmn)
	switch tnmn {
	case "nepenthes" :
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Asli Indonesia")
	case "venus" :
		fmt.Println("Termasuk Tanaman Karnivora")
		fmt.Println("Bukan Asli Indonesia")
	default :
		fmt.Print("Tidak termasuk Tanaman Karnivora")
	}
}
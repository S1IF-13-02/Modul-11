package main

import "fmt"

func main() {

	var n string
	fmt.Print("nama tanaman : ")
	fmt.Scan(&n)
	switch n {
	case "nepenthes", "kantong semar", "rafflesia":
		fmt.Println("termasuk tanaman karnivora ")
		fmt.Println("dan asli dari indonesia")

	case "venus":
		fmt.Println("termasuk tanaman karnivora ")
		fmt.Println("dan bukan asli dari indonesia")

	default:
		fmt.Println("bukan tanaman karnivora")

	}
}

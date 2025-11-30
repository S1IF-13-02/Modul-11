package main 

import "fmt"

func main (){
	var jam12,jam24 int
	var ampm string
	fmt.Scan(&jam24)
	switch {
	case jam24 == 0 :
		jam12 = 12 
		ampm = "AM"
	case jam24 < 12 :
		jam12 = jam24
		ampm = "AM"
	case jam24 == 12 :
		jam12 = 12 
		ampm = "PM"
	case jam24 > 12 :
		jam12 = jam24 - 12
		ampm = "PM"
	}
	fmt.Println(jam12, ampm)
}
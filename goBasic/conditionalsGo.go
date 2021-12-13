package main

import "fmt"

var stats bool

func main() {
	stats = true
	if stats == true {
		// if stats = false; stats==true{
		fmt.Println(stats)
		// } else if stats == false {
	} else {
		fmt.Println("It's false")
	}

	switch number := 7; number {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("Bigger than 5")
	}
}

package main

import "fmt"

var stats bool

func main() {
	stats = true
	if stats == true {
		// if stats = false; stats==true{
		fmt.Println(stats)
	} else {
		fmt.Println("It's false")
	}
}

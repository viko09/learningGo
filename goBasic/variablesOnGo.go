package main

import "fmt"

// Variables that can be used in any function
// var number int
// var text string
var status = true

func main() {
	// Variables that can be used only in this function
	// var number1 int
	// number2 := 3
	// var numberX, numberY, numberZ int
	numberX, numberY, numberZ, text1, status := 3.4, 5, 8, "Victor", false

	// var coin float32 = 0
	// number2 = coin  <--- Error

	fmt.Println(numberX)
	fmt.Println(numberY)
	fmt.Println(numberZ)
	fmt.Println(text1)
	fmt.Println(status)
	showStats()

	// Converting data
	newNum := 5
	text := "Texto"

	var coin int64 = 0
	newNum = int(coin)

	text = fmt.Sprintf("%d", coin)
	fmt.Println(newNum)
	fmt.Println(text)
}

func showStats() {
	fmt.Println(status)
}

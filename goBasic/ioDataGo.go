package main

import (
	"bufio"
	"fmt"
	"os"
)

var number1 int
var number2 int
var result int
var legend string

func main() {
	fmt.Println("Please insert your first number: ")
	fmt.Scanf("%d", &number1)

	fmt.Println("Please insert your second number: ")
	fmt.Scanf("%d", &number2)

	fmt.Println("%d", number1+number2)

	fmt.Println("Description: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		// fmt.Println("%d", number2 - number1)
		legend = scanner.Text()
	}

	result = number2 - number1
	fmt.Println(legend, result)
}

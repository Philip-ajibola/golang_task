package main

import "fmt"

func main() {
	firstNumber := 0
	secondNumber := 0
	for {
		fmt.Print("Enter first number: ")
		_, err := fmt.Scanf("%d", &firstNumber)
		if err != nil {
			fmt.Println("Invalid input")
			var invalidInput string
			fmt.Scanln(invalidInput)
		} else {
			break
		}
	}
	for {
		fmt.Print("Enter first number: ")
		_, err := fmt.Scanf("%d", &secondNumber)
		if err != nil {
			fmt.Println("Invalid input")
			var invalidInput string
			fmt.Scanln(invalidInput)
		} else {
			break
		}
	}
	fmt.Print(checkNumberPrecedence(firstNumber, secondNumber))
}

func checkNumberPrecedence(num1, num2 int) int {
	if num1 > num2 {
		return 1
	} else if num1 < num2 {
		return -1
	} else {
		return 0
	}
}

package main

import "fmt"

func main() {
	number := 0
	for {
		fmt.Print("Enter your number: ")
		_, err := fmt.Scanf("%d", &number)
		if err != nil {
			fmt.Print("Invalid Input")
			var invalidInput string
			fmt.Scanln(&invalidInput)
		} else {
			break
		}
	}
	fmt.Printf("The Factorial of %d is %d", number, factorial(number))
}
func factorial(num int) int {
	for counter := num - 1; counter >= 1; counter-- {
		num *= counter
	}
	return num
}

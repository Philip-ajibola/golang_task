package main

import (
	"fmt"
	"strings"
)

func main() {
	var numbers [5]int
	number := 0
	fmt.Println("Enter Number Between 1 and 30: ")
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &number)
		for number < 1 || number > 30 {
			fmt.Println("please enter a number between 1 and 30")
			fmt.Println("Enter Number Between 1 and 30: ")
			var invalidInput string
			fmt.Scanln(&invalidInput)
			fmt.Scanf("%d", &number)
		}
		numbers[i] = number
	}
	for counter := 0; counter < 5; counter++ {
		fmt.Println(strings.Repeat("*", numbers[counter]))
	}
}

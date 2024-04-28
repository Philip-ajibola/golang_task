package main

import "fmt"

func main() {
	number := 0
	fmt.Print("Enter An Odd Number from 1-19: ")
	fmt.Scanf("%d", &number)
	number1 := 0
	for number%2 == 0 {
		fmt.Print("Enter Am Odd Number From 1-19,Listen To Instructions ")
		fmt.Scanf("%d", &number)
	}
	number1 = (number / 2) + 1
	for count := 1; count <= number1; count++ {
		for counter := number1; counter >= count; counter-- {
			fmt.Print(" ")
		}
		for count1 := 1; count1 <= count; count1++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	for count2 := 1; count2 <= number1; count2++ {
		fmt.Print(" ")
		for count3 := 1; count3 <= count2; count3++ {
			fmt.Print(" ")
		}
		for count4 := number1 - 1; count4 >= count2; count4-- {
			fmt.Print("* ")
		}
		fmt.Println()
	}

}

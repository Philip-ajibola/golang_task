package main

import "fmt"

func main() {
	numberOfInput := 0
	number := 0
	highest := 0
	fmt.Print("How many number do you want to input : ")
	fmt.Scanf("%d", &numberOfInput)
	smallest := 0
	fmt.Print("Enter numbers : ")
	fmt.Scanf("%d", &smallest)
	for counter := 1; counter < numberOfInput; counter++ {
		fmt.Scanf("%d", &number)
		if number > highest {
			highest = number
		}
		if number < smallest {
			smallest = number
		}
	}
	fmt.Println("The highest number :", highest)
	fmt.Println("The smallest number :", smallest)

	fmt.Println("Sum Of the number is :", highest+smallest)
}

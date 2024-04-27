package main

import "fmt"

func main() {
	largestNumber := 0
	secondLargestNumber := 0
	number := 0
	for counter := 0; counter < 10; counter++ {
		for {
			fmt.Print("enter Integer ")
			_, err := fmt.Scanf("%d", &number)
			if err == nil {
				break
			} else {
				fmt.Println("Error: Please Enter A Valid Number ")
				var invalidInput string
				fmt.Scanln(&invalidInput)
			}
		}
		if number > largestNumber {
			secondLargestNumber = largestNumber
			largestNumber = number
		}
	}
	fmt.Println()
	fmt.Println("Largest Number: ", largestNumber)
	fmt.Println("Second Largest Number: ", secondLargestNumber)
}

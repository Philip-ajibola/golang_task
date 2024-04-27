package main

import "fmt"

var numberOfItemSold int
var valueOfEachItemSold float32
var totalValueOfItemSold float32

func main() {
	for i := 1; i <= 4; i++ {
		for {
			fmt.Print("Enter Number of Item Sold: ")
			_, err := fmt.Scanf("%d", &numberOfItemSold)
			if err == nil {
				break
			} else {
				fmt.Println("Invalid Input")
				var invalidInput string
				fmt.Scanln(&invalidInput)
			}
		}
		for {
			fmt.Print("Enter Value of Item Sold: ")
			_, err := fmt.Scanf("%f", &valueOfEachItemSold)
			if err == nil {
				break
			} else {
				fmt.Println("Invalid Input ")
				var invalidInput string
				fmt.Scanln(&invalidInput)
			}
		}
		totalValueOfItemSold = float32(numberOfItemSold) * valueOfEachItemSold
		fmt.Println("Your total wage for this week is: ", totalValueOfItemSold*(9/100))
	}
}

package main

import "fmt"

func main() {
	firstNumber := 0
	for {
		fmt.Print("Enter  number: ")
		_, err := fmt.Scanf("%d", &firstNumber)
		if err != nil {
			fmt.Println("Enter A valid Input.")
			var invalidInput string
			fmt.Scanln(&invalidInput)
		} else {
			break
		}
	}

	numbers := 0
	for numbers < firstNumber {
		for {
			fmt.Print("Enter number: ")
			_, err := fmt.Scanf("%d", &numbers)
			if err != nil {
				fmt.Println("Enter A valid Input.")
				var invalidInput string
				fmt.Scanln(&invalidInput)
			} else {
				break
			}
		}
	}

}

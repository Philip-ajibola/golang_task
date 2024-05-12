package main

import "fmt"

func main() {
	number := 0
	exponential := 1.0
	for {
		fmt.Println("Enter A Number: ")
		_, err := fmt.Scanf("%d", &number)
		if err != nil {
			fmt.Print("Invalid Input")
			var invalidInput string
			fmt.Scanln(&invalidInput)
		} else {
			break
		}
	}
	for counter := 1; counter <= number; counter++ {
		exponential += 1.0 / float64(factorial1(counter))
	}
	fmt.Println("Exponentital is ", exponential)
}
func factorial1(num int) int {
	for counter := num - 1; counter >= 1; counter-- {
		num *= counter
	}
	return num
}

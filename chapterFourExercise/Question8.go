package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := 0
	convertedNumber := ""
	for {
		fmt.Print("Enter an Integer of Length 5")
		_, err := fmt.Scanln(&number)
		convertedNumber = strconv.Itoa(number)
		if err != nil {
			fmt.Print("Invalid Input")
			var invalidInput string
			fmt.Scanln(&invalidInput)
		} else {
			break
		}
	}
	if len(convertedNumber) != 5 {
		fmt.Println("Number is not valid")
		var invalidInput string
		fmt.Scanln(invalidInput)
		main()
	}
	fmt.Print(checkIfNumberIsPalindrome(convertedNumber))
}
func checkIfNumberIsPalindrome(number string) bool {
	splittedNumber := []rune(number)
	counter := len(splittedNumber) - 1
	count := 0
	for count < counter {
		splittedNumber[count], splittedNumber[counter] = splittedNumber[counter], splittedNumber[count]
		counter--
		count++
	}
	return string(splittedNumber) == number
}

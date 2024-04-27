package main

import (
	"fmt"
)

func main() {
	triangleLength := 0
	for {
		fmt.Print("Enter the length of the triangle you want: ")
		_, err := fmt.Scanf("%d", &triangleLength)
		if err != nil || triangleLength <= 0 {
			fmt.Println("Error: Invalid input. Please enter a valid positive integer.")
			var invalidInput string
			fmt.Scanln(&invalidInput)
		} else {
			break
		}
	}
	for counter := 0; counter < triangleLength; counter++ {
		fmt.Println(repeat("*", counter))
	}
}
func repeat(asterisk string, number int) string {
	result := ""
	for counter := 0; counter < number; counter++ {
		result += " " + asterisk
	}
	return result
}

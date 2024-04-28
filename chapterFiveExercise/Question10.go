package main

import "fmt"

func main() {

	for count := 1; count <= 10; count++ {
		for counter := 10; counter >= count; counter-- {
			fmt.Print(" ")
		}
		for count1 := 1; count1 <= count; count1++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	for count2 := 1; count2 <= 10; count2++ {
		fmt.Print(" ")
		for count3 := 1; count3 <= count2; count3++ {
			fmt.Print(" ")
		}
		for count4 := 9; count4 >= count2; count4-- {
			fmt.Print("* ")

		}
		fmt.Println()
	}
}

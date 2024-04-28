package main

import (
	"fmt"
)

func main() {
	spaceCounter := 1
	fmt.Printf("%s%15s%15s%23s\n", "(A)", "(B)", "(C)", "(D)")
	for count := 1; count <= 10; count++ {
		for counter := 1; counter <= count; counter++ {
			fmt.Print("*")
		}
		for counter := 15; counter >= count; counter-- {
			fmt.Print(" ")
		}
		for counter1 := 10; counter1 >= count; counter1-- {
			fmt.Print("*")
		}

		for counter := 5 + spaceCounter; counter >= 1; counter-- {
			fmt.Print(" ")
		}

		for counter3 := 10; counter3 >= count; counter3-- {
			fmt.Print("*")
		}
		for counter := 15; counter >= count; counter-- {
			fmt.Print(" ")
		}
		for counter := 1; counter <= count; counter++ {
			fmt.Print("*")
		}
		fmt.Println()
		spaceCounter += 2
	}
}

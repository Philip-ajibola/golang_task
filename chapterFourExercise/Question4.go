package main

import "fmt"

func main() {
	highest := 0
	number := 0
	for counter := 0; counter < 10; counter++ {
		fmt.Printf("Enter Integer value %d\n", counter+1)
		fmt.Scanf("%d", &number)

		if number > highest {
			highest = number
		}
		fmt.Printf("Highest number is %d\n", highest)
	}
}

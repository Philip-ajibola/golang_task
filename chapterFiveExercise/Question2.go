package main

import (
	"fmt"
)

func main() {
	total := 0
	for counter := 1; counter <= 30; counter++ {
		if counter%3 == 0 {
			total += counter
		}
	}
	fmt.Println("The Total  The Sum Of Number Divisible By 3 Between 1 - 30 is: ", total)
}

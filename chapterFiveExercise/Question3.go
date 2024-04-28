package main

import "fmt"

func main() {
	total := 0
	for counter := 1; counter <= 100; counter++ {
		fmt.Printf("%d + %d = %d\n", total, counter, total+counter)
		total += counter
	}
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var principal float64 = 1000
	var rate float64 = 0.05

	for rate < 0.10 {
		calculateInterest(rate, principal)
		rate += 0.01
	}

}
func calculateInterest(interest, principal float64) {
	fmt.Printf("%s%20s\n", "Year", "Amount on deposit")

	var amount float64
	for year := 1; year < 10; year++ {
		amount = principal * math.Pow(1+interest, float64(year))
		fmt.Printf("%4d%20.2f\n", year, amount)
	}
	fmt.Println()

}

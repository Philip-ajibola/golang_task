package main

import (
	"fmt"
	"math"
)

func main() {
	var principal float64 = 1000
	var rate float64 = 0.05

	fmt.Printf("%-10s%20s%27s\n", "Year", "Amount On Deposit($ dollars)", "Amount Deposit(cents)")

	for year := 1; year <= 12; year++ {
		amount := int((principal * math.Pow(1+rate, float64(year))) * 100)
		dollar := amount / 100
		cents := amount % 100
		fmt.Printf("%4d%20d$%27dcent\n", year, dollar, cents)

	}
}

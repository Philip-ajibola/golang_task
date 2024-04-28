package main

import "fmt"

func main() {
	number := 0
	total := 1.0
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
	for count := 1; count < number; count++ {
		total += float64(pow1(number, count) / factorial2(count))
	}
}
func factorial2(num int) int {
	for counter := num - 1; counter >= 1; counter-- {
		num *= counter
	}
	return num
}
func pow1(num1, num2 int) int {
	num3 := num1
	for counter := 1; counter < num2; counter++ {
		num3 *= num1
	}
	fmt.Println(num1)
	return num3
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := 0
	convertedNumber := ""
	for {
		fmt.Print("Enter a number: ")
		_, err := fmt.Scanf("%d", &number)
		if err != nil {
			return
		} else {
			break
		}
	}
	convertedNumber = strconv.Itoa(number)
	fmt.Println("The Decimal Equivalent Of your input is; ", reverseString(convertedNumber))
}

func reverseString(number string) int {
	splittedNumber := []rune(number)
	counter := len(splittedNumber) - 1
	count := 0
	for count < counter {
		splittedNumber[count], splittedNumber[counter] = splittedNumber[counter], splittedNumber[count]
		counter--
		count++
	}
	fmt.Print(splittedNumber)
	intArray := make([]int, len(splittedNumber))
	for count, runeValue := range splittedNumber {
		str := string(runeValue)
		intValue, _ := strconv.Atoi(str)
		intArray[count] = intValue
	}
	return giveDecimalEquivalentOf(intArray)

}
func giveDecimalEquivalentOf(binary []int) int {
	number := len(binary)
	total := binary[0] * 1
	for counter := 1; counter < number; counter++ {
		total += binary[counter] * pow(2, counter)
	}
	return total
}

func pow(num1, num2 int) int {
	num3 := num1
	for counter := 1; counter < num2; counter++ {
		num3 *= num1
	}
	fmt.Println(num1)
	return num3
}

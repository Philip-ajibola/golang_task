package main

import (
	"fmt"
	"strings"
)

var milePerGallon float32
var gallonUsed int
var mileCovered int
var totalMilePerGallon float32

func main() {
	totalMilePerGallon = calculateMilePerGallon()
	fmt.Println("Total Mile Per Gallon For All trip is ", totalMilePerGallon)
}
func calculateMilePerGallon() float32 {
	return displayTotalMilePerGallon()
}

func displayTotalMilePerGallon() float32 {
	for {
		mileCovered = collectMiles()
		gallonUsed = collectGallon()
		totalMilePerGallon = calculateTotalMilePerGallon()
		var sentinel string
		for {
			fmt.Print("Do You Want to Enter For More Trip (yes or no) ")
			_, err := fmt.Scanln(&sentinel)
			if err != nil {
				fmt.Print("Please enter a valid input ")
				var invalidInput string
				fmt.Scanln(&invalidInput)
			} else {
				break
			}
		}
		if strings.EqualFold(sentinel, "no") {
			break
		}
	}
	return totalMilePerGallon
}

func calculateTotalMilePerGallon() float32 {
	milePerGallon = float32(mileCovered) / float32(gallonUsed)
	totalMilePerGallon += milePerGallon
	fmt.Println("MillePerGallon For This Trip is ", milePerGallon)
	return totalMilePerGallon
}

func collectGallon() int {
	for {
		fmt.Print("Enter Gallon used ")
		_, err := fmt.Scanf("%d", &gallonUsed)
		if err != nil {
			fmt.Print("Invalid Input ")
			var invalidInput string
			fmt.Scanln(&invalidInput)
		} else {
			break
		}
	}
	return gallonUsed
}

func collectMiles() int {
	for {
		fmt.Print("Enter Mile: ")
		_, err := fmt.Scanf("%d", &mileCovered)
		if err != nil {
			fmt.Print("Enter Mile: ")
			var invalidInput string
			fmt.Scanln(&invalidInput)
		} else {
			break
		}
	}
	return mileCovered
}

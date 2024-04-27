package main

import (
	"fmt"
)

func main() {
	var x1, y1 float64
	fmt.Println("Enter the coordinates of the first point:")
	for {
		fmt.Print("x1: ")
		_, err := fmt.Scanf("%f", &x1)
		if err != nil {
			fmt.Print("Invalid input\n")
			var invalid string
			fmt.Scanln(invalid)
		} else {
			break
		}

	}
	for {
		fmt.Print("y1: ")
		_, err := fmt.Scan(&y1)
		if err != nil {
			fmt.Print("Invalid input\n")
			var invalid string
			fmt.Scanln(invalid)
		} else {
			break
		}

	}

	var x2, y2 float64
	fmt.Println("\nEnter the coordinates of the second point:")
	for {
		fmt.Print("x2: ")
		_, err := fmt.Scan(&x2)
		if err != nil {
			fmt.Print("Invalid input\n")
			var invalid string
			fmt.Scanln(invalid)
		} else {
			break
		}

	}
	for {
		fmt.Print("y2: ")
		_, err := fmt.Scan(&y2)
		if err != nil {
			fmt.Print("Invalid input\n")
			var invalid string
			fmt.Scanln(invalid)
		} else {
			break
		}
	}

	if (x1 == x2 && y1 != y2) || (x1 != x2 && y1 == y2) {
		fmt.Println("\nThe points are located on a line perpendicular to an axis.")
	} else {
		fmt.Println("\nThe points are not located on a line perpendicular to an axis.")
	}
}

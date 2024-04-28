package main

import (
	"fmt"
	"strings"
)

var countA int
var countB int
var countC int
var countD int

func main() {
	studentName := ""
	studentGrade := ""
	for counter := 0; counter < 5; counter++ {
		fmt.Print("What is your name? ")
		fmt.Scan(&studentName)
		fmt.Print("What is your grade? ")
		fmt.Scan(&studentGrade)
		for len(studentGrade) != 1 {
			fmt.Print("Invalid grade Inputted\nInput A valid grade")
			fmt.Scan(&studentGrade)
		}
		for !strings.Contains("ABCD", studentGrade) {
			fmt.Print("Invalid grade\nInput A valid grade")
			fmt.Scan(&studentGrade)
		}
		countGrades(studentGrade)
	}
	fmt.Println("Number of student with grade A is: ", countA)
	fmt.Println("Number of student with grade B is: ", countB)
	fmt.Println("Number of student with grade C is: ", countC)
	fmt.Println("Number of student with grade D is: ", countD)
}
func countGrades(alpha string) {
	switch alpha {
	case "A":
		countA++
	case "B":
		countB++
	case "C":
		countC++
	case "D":
		countD++
	default:
		fmt.Println("Invalid grade")
	}

}

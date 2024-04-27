package main

import (
	"fmt"
	"strings"
)

var accountNumber int
var balanceAtMonthBeginning int
var totalOfAllItemCharged int
var totalCreditApplied int
var creditLimit int
var newBalance int

func main() {
	for {
		accountNumber = collectAccountNumber()
		balanceAtMonthBeginning = collectBalance()
		totalOfAllItemCharged = collectItemCharged()
		totalCreditApplied = collectAllCreditApplied()
		creditLimit = collectCreditLimit()
		newBalance = balanceAtMonthBeginning + totalOfAllItemCharged - totalCreditApplied
		if newBalance < creditLimit {
			fmt.Println("You Already Exit Your Credit limit ")
		}
		var sentinel string
		sentinel = breakLoop(sentinel)
		if strings.EqualFold(sentinel, "no") {
			break
		}
	}
}

func breakLoop(sentinel string) string {
	for {
		fmt.Print("More Workers ? (Yes or No) ")
		_, err := fmt.Scanln(&sentinel)
		for !strings.EqualFold(sentinel, "yes") || !strings.EqualFold(sentinel, "no") {
			fmt.Println("Please Enter a valid Input: (Yes or No) ")
			_, err = fmt.Scanln(&sentinel)
		}
		if err != nil {
			fmt.Println("Please Enter A Valid Response: (Yes or No ) ")

		} else {
			break
		}
	}

	return sentinel
}

func collectAccountNumber() int {
	for {
		fmt.Print("Enter Account Number: ")
		_, err := fmt.Scanf("%d", &accountNumber)
		if err != nil {
			fmt.Print("Invalid input\nEnter Account Number: ")
			var invalidInput string
			fmt.Scanln(&invalidInput)
		} else {
			break
		}
	}
	return accountNumber
}

func collectCreditLimit() int {
	for {
		fmt.Print("Enter Your Credit Limit: ")
		_, err := fmt.Scanf("%d", &creditLimit)
		if err != nil {
			fmt.Print("Invalid input\nEnter Your Credit Limit: ")
			var invalidInput string
			fmt.Scanln(&invalidInput)
		} else {
			break
		}
	}
	return creditLimit
}

func collectAllCreditApplied() int {
	for {
		fmt.Print("Enter Total Credit Applied: ")
		_, err := fmt.Scanf("%d", &totalCreditApplied)
		if err != nil {
			fmt.Print("Invalid input\nEnter Total Credit Applied: ")
			var invalidInput string
			fmt.Scanln(&invalidInput)
		} else {
			break
		}
	}
	return totalCreditApplied
}

func collectItemCharged() int {
	for {
		fmt.Print("Enter Total Of All Item Charged: ")
		_, err := fmt.Scanf("%d", &totalOfAllItemCharged)
		if err != nil {
			fmt.Print("Invalid input\nEnter Total Of All Item Charged: ")
			var invalidInput string
			fmt.Scanln(&invalidInput)
		} else {
			break
		}
	}
	return totalOfAllItemCharged
}

func collectBalance() int {
	for {
		fmt.Print("Enter Balance: ")
		_, err := fmt.Scanf("%d", &balanceAtMonthBeginning)
		if err != nil {
			fmt.Print("Invalid input\nEnter Account Number: ")
			var invalidInput string
			fmt.Scanln(&invalidInput)
		} else {
			break
		}
	}

	return balanceAtMonthBeginning
}

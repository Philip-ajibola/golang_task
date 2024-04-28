package main

import "fmt"

var accountNumber int
var state string
var makeAndModel string

func main() {
	fmt.Println("Enter Account Number")
	fmt.Scanf("%d", &accountNumber)
	fmt.Println("Enter States Code E.g(New York(NY))")
	fmt.Scanf("%s", &state)
	validateState(state)
	fmt.Println("Enter MakeAndModel E.g(Toyota Camry)")
	fmt.Scanf("%s", &makeAndModel)

	condition := isNoFaultState()
	if condition {
		fmt.Println(state, "is  A No fault state")
	} else {
		fmt.Println(state, "is not No fault state")
	}

}

func validateState(stateCode string) {
	switch stateCode {
	case "NJ", "NY", "MA", "PA", "CT", "VT", "NH", "ME":
	default:
		fmt.Println("State Code is not valid")
	}

}
func isNoFaultState() bool {
	var noFaultState bool
	switch state {
	case "MA":
		noFaultState = true
	case "NJ":
		noFaultState = true
	case "NY":
		noFaultState = true
	case "PA":
		noFaultState = true

	default:
		noFaultState = false
	}
	return noFaultState
}

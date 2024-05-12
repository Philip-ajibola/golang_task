package CheckOutList

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var numberOfItems []int
var priceOFEachItem []float64
var nameOfItems []string
var totalPricePerItem []float64
var customerNames string
var cashierName string
var discount int
var billTotal float64

func CheckOutList() {
	collectAllInput()
	displayInvoice()
	displayReceipt()
}
func collectAllInput() {
	collectBuyersName()
	for {
		collectItemsName()
		unit := collectNumberOfItemBought()
		pricePerUnit := CollectPriceOfItem()
		totalPricePerItem = append(totalPricePerItem, pricePerUnit*float64(unit))
		userResponse := collectUserResponse()
		if strings.EqualFold(userResponse, "No") {
			break
		}
	}
}

func collectUserResponse() string {
	userResponse := ""
	fmt.Print("Add More Item ")
	fmt.Scan(&userResponse)
	for !strings.EqualFold(userResponse, "Yes") && !strings.EqualFold(userResponse, "No") {
		fmt.Print("Invalid Input")
		var invalidInput string
		fmt.Scan(&invalidInput)
		fmt.Print("Add More Item ")
		fmt.Scan(&userResponse)
	}
	return userResponse
}

func CollectPriceOfItem() float64 {
	pricePerUnit := 0.0
	fmt.Print("What's the price per unit ")
	fmt.Scanf("%f", &pricePerUnit)
	for pricePerUnit <= 0 {
		fmt.Print("Please Enter Valid Price")
		var invalidInput string
		fmt.Scan(&invalidInput)
		fmt.Print("What's the price per unit ")
		fmt.Scanf("%f", &pricePerUnit)
	}
	priceOFEachItem = append(priceOFEachItem, pricePerUnit)
	return pricePerUnit
}

func collectNumberOfItemBought() int {
	unit := 0
	fmt.Print("How many did user Buy ")
	fmt.Scan(&unit)
	for unit <= 0 {
		fmt.Print("Please Enter Valid Unit ")
		var invalidInput string
		fmt.Scan(&invalidInput)
		fmt.Print("How many did user Buy ")
		fmt.Scan(&unit)
	}
	numberOfItems = append(numberOfItems, unit)
	return unit
}

func collectItemsName() {
	itemBought := ""
	fmt.Print("what did the user buy ??? ")
	fmt.Scan(&itemBought)
	for len(itemBought) == 0 {
		fmt.Print("Please Enter an Item Name")
		var invalidInput string
		fmt.Scan(&invalidInput)
		fmt.Print("what did the user buy ??? ")
		fmt.Scan(&itemBought)
	}
	nameOfItems = append(nameOfItems, itemBought)
}

func collectBuyersName() {
	pattern := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	fmt.Print("What's the customers name ")
	fmt.Scan(&customerNames)
	for !pattern.MatchString(strings.ToLower(customerNames)) {
		fmt.Print("Please Enter A VaLid Name")
		var invalidInput string
		fmt.Scanln(&invalidInput)
		fmt.Print("What's the customers name ")
		fmt.Scan(&customerNames)
	}
}

func collectCashierName() {
	fmt.Print("What is Your Name ")
	fmt.Scan(&cashierName)
	for len(cashierName) == 0 {
		fmt.Print("Invalid Input ")
		var invalidInput string
		fmt.Scan(&invalidInput)
		fmt.Print("What is Your Name ")
		fmt.Scan(&cashierName)
	}
	fmt.Print("How much discount will customer get ")
	fmt.Scan(&discount)
	displayHeader()
}

func displayHeader() {
	currentTime := time.Now()
	formattedDate := currentTime.Format("02-DEC-2006 03:04 PM")
	fmt.Printf(`
	SEMICOLON STORE
	MAIN BRANCH 
	LOCATION: 312, HEBERT MACAULAY WAY, SABO YABA, LAGOS.
	TEL: 03293828343
	Date: %s
	Cashier: %s
	Customer-Name: %s
`, formattedDate, cashierName, customerNames)
}

func displayInvoice() {
	collectCashierName()
	fmt.Printf(`
	============================================================================
	%-20s\t%-20s\t%-20s\t%-20s                         
	----------------------------------------------------------------------------
`, "ITEM", "QTY", "Price", "TotalAmount")

	for counter := 0; counter < len(numberOfItems); counter++ {
		fmt.Printf("%-20s\t%-20d\t%-20.2f\t%-20.2f", nameOfItems[counter], numberOfItems[counter], priceOFEachItem[counter], totalPricePerItem[counter])
		fmt.Println()
	}
	total := 0.0
	for counter := 0; counter < len(numberOfItems); counter++ {
		total += totalPricePerItem[counter]
	}
	discountAmount := total * (float64(discount) / 100.0)
	valueAddedTax := total * (17.5 / 100.0)
	billTotal = total - discountAmount + valueAddedTax

	fmt.Printf(`
	-----------------------------------------------------------------------------
										SubTotal:             %.2f 
										Discount:             %.2f
									  VAT @ 17.50%s           %.2f
	------------------------------------------------------------------------------
									Bill Total:               %.2f
	==============================================================================
		THIS IS NOT A RECIEPT KINDLY PAY      %.2f
	==============================================================================
`, total, discountAmount, "%", valueAddedTax, billTotal, billTotal)
}
func displayReceipt() {
	amountPaid := 0.0
	fmt.Print("How Much Did The User Pay ")
	fmt.Scanf("%f", &amountPaid)
	for amountPaid < billTotal {
		fmt.Print("Money Not Complete")
		var invalidInput string
		fmt.Scan(&invalidInput)
		fmt.Print("How Much Did The User Pay ")
		fmt.Scanf("%f", &amountPaid)
	}
	displayHeader()
	fmt.Println()
	fmt.Printf(`
	============================================================================
	%-20s\t%-20s\t%-20s\t%-20s                         
	----------------------------------------------------------------------------
`, "ITEM", "QTY", "Price", "TotalAmount")

	for counter := 0; counter < len(numberOfItems); counter++ {
		fmt.Printf("     %-20s\t%-20d\t%-20.2f\t%-20.2f", nameOfItems[counter], numberOfItems[counter], priceOFEachItem[counter], totalPricePerItem[counter])
		fmt.Println()
	}
	total := 0.0
	for counter := 0; counter < len(numberOfItems); counter++ {
		total += totalPricePerItem[counter]
	}
	discountAmount := total * (float64(discount) / 100.0)
	valueAddedTax := total * (17.5 / 100.0)
	billTotal = total - discountAmount + valueAddedTax

	fmt.Printf(`
	-----------------------------------------------------------------------------
										SubTotal:             %.2f 
										Discount:             %.2f
									  VAT @ 17.50%s           %.2f
	------------------------------------------------------------------------------
									Bill Total:               %.2f
									AmountPaid:               %.2f
									Balance:                  %.2f
	==============================================================================
		Thank for the patronage  :)
	==============================================================================
`, total, discountAmount, "%", valueAddedTax, billTotal, billTotal)
}

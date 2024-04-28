package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter Number 1-12: ")
	fmt.Scanln(&number)

	switch number {
	case 1:
		fmt.Println("On The First Day Of Christmas\nMy true love sent to me:")
	case 2:
		fmt.Println("On The Second Day Of Christmas\nMy true love sent to me:")
	case 3:
		fmt.Println("On The Third Day Of Christmas\nMy true love sent to me:")
	case 4:
		fmt.Println("On The Fourth Day Of Christmas\nMy true love sent to me:")
	case 5:
		fmt.Println("On The Fifth Day Of Christmas\nMy true love sent to me:")
	case 6:
		fmt.Println("On The Sixth Day Of Christmas\nMy true love sent to me:")
	case 7:
		fmt.Println("On The Seventh Day Of Christmas\nMy true love sent to me:")
	case 8:
		fmt.Println("On The Eighth Day Of Christmas\nMy true love sent to me:")
	case 9:
		fmt.Println("On The Ninth Day Of Christmas\nMy true love sent to me:")
	case 10:
		fmt.Println("On The Tenth Day Of Christmas\nMy true love sent to me:")
	case 11:
		fmt.Println("On The Eleventh Day Of Christmas\nMy true love sent to me:")
	case 12:
		fmt.Println("On The Twelfth Day Of Christmas\nMy true love sent to me:")
	}

	switch number {
	case 12:
		fmt.Println("Twelve drummers drumming")
		fallthrough
	case 11:
		fmt.Println("Eleven pipers piping")
		fallthrough
	case 10:
		fmt.Println("Ten lords a-leaping")
		fallthrough
	case 9:
		fmt.Println("Nine ladies dancing")
		fallthrough
	case 8:
		fmt.Println("Eight maids a-milking")
		fallthrough
	case 7:
		fmt.Println("Seven swans a-swimming")
		fallthrough
	case 6:
		fmt.Println("Six geese a-laying")
		fallthrough
	case 5:
		fmt.Println("Five golden rings")
		fallthrough
	case 4:
		fmt.Println("Four calling birds ")
		fallthrough
	case 3:
		fmt.Println("Three French hens")
		fallthrough
	case 2:
		fmt.Println("Two turtle doves and A ")
		fallthrough
	case 1:
		fmt.Println("A partridge in a pear tree ")
	}
}

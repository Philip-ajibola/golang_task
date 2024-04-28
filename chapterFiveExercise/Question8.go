package main

import "fmt"

func main() {
	for side1 := 1; side1 <= 500; side1++ {
		for side2 := 1; side2 <= 500; side2++ {
			for hypothenus := 1; hypothenus <= 500; hypothenus++ {
				if (side1*side1)+(side2*side2) == (hypothenus * hypothenus) {
					fmt.Printf("%4d\t%4d\t%10d\n", side1, side2, hypothenus)
					break
				}
				if (hypothenus * hypothenus) > (side1*side1)+(side2*side2) {
					break
				}

			}
		}
	}
}

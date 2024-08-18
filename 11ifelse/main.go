package main

import "fmt"

func main() {
	fmt.Println("Wellcome to ifelse in golang")

	x := 10
	if x > 5 { //if else ladder
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than or equal to 5")
	}

	if x < 5 { //if -else if -else ladder
		fmt.Println("x is less than 5")
	} else if x == 10 {
		fmt.Println("x is equal to 10")
	} else {
		fmt.Println("x is greater than 5 but not 10")
	}
	//variables can decalre in if statements
	if y := 20; y > x {
		fmt.Println("y is greater than x")
	} else if y == x {
		fmt.Println("y is equal to x")
	} else {
		fmt.Println("y is less than x")
	}

	loginCount := 23
	var message string
	if loginCount < 10 {
		message = "Regular users"
	} else if loginCount > 10 {
		message = "Watch out"
	} else {
		message = "Exactly 10 login count"
	}

	fmt.Println("message:", message)

	if x := 9; x%2 == 0 {
		fmt.Println("x is even")
	} else {
		fmt.Println("x is odd")
	}
}

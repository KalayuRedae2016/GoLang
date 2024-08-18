package main

import "fmt"

func main() {
	fmt.Println("wellcome to swich in golang")

	day := "monday"
	switch day {

	case "monday":
		fmt.Println("first day of work week")
	case "friday":
		fmt.Println("last day of work week")
	case "staturday", "sunday":
		fmt.Println("it is weekend")
	default:
		fmt.Println("midweek days")
	}

	//switch can use with out expression for multiple cases
	num := 7
	switch {
	case num > 7:
		fmt.Println("greater than seven")
	case num == 7:
		fmt.Println("equal seven")
	case num < 7:
		fmt.Println("less than seven")

	}

	//switch with fallthrough-force next block to execute
	n := 2
	switch n {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
		fallthrough
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("Unknown number")
	}

}

package main

import "fmt"

var LoggedTokken = "2423gege" //Public variable and no need to be used outside functions

func main() {
	var userName string = "kalayu"
	fmt.Println("Hello", userName)
	fmt.Printf("variable type is :%T\n", userName)

	var isLogged bool = false
	fmt.Println("Hello", isLogged)
	fmt.Printf("variable type is :%T\n", isLogged)

	var anottherVariable int
	fmt.Println(anottherVariable)

	var website = "www.kaleb.com"
	fmt.Println(website)

	//shorthand
	numberOfUsers := 54334
	website2 := "www.kaleb.com"
	fmt.Println("number of users", numberOfUsers)
	fmt.Println(website2)

	fmt.Println("LoggedToken", LoggedTokken)

}

package main

import "fmt"

func main() {
	fmt.Println("Wellcome to slices in golang")
	var fruits = []string{} //nill slices or empty slices
	fmt.Println("Fruits list", fruits)

	fruits = append(fruits, "Mango", "bananna", "organge") //append to slices
	fmt.Println("Fruits list", fruits)

	fmt.Println("length of fruits", len(fruits))     //length of slices
	fmt.Println("capacity of fruits", cap((fruits))) //capacity of slices

}

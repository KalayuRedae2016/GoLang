package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string //declare array with 4 elements of string types
	fruitList[0] = "Apple"
	fruitList[1] = "tomatto"
	fruitList[3] = "Peach"
	//fruitList[4] = "kaleb" index of out range

	fmt.Println("FruitList ", fruitList)           //displaying the array
	fmt.Println("fruit 1", fruitList[0])           ///accessing single array using index
	fmt.Println("Length of fruit", len(fruitList)) //length of array

	var vegList = [3]string{"potato", "carrot", "beans"} //declaring and intialize array
	fmt.Println("VegList ", vegList)
	fmt.Println("len of veglist", len(vegList))

	//modifing arrays
	fruitList[2] = "Wheat"
	fmt.Println("fruitlist", fruitList)

	//iterating arrays
	for i := 0; i < len(fruitList); i++ {
		fmt.Println(fruitList[i])
	}
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Wellcome to maps in golang")
	userScores := make(map[string]int)
	userScores["hailom"] = 95
	userScores["kalayu"] = 64
	userScores["berhe"] = 65
	userScores["tadese"] = 90
	fmt.Println("userscores", userScores)

	//Read
	berheScore := userScores["berhe"]
	kalayuScore := userScores["kalayu"]
	fmt.Println("Berhe's score", berheScore)
	fmt.Println("Kalayu's score", kalayuScore)

	//Update
	userScores["tadese"] = 92
	fmt.Println("Updated tadese's score", userScores["tadese"])

	//delete
	delete(userScores, "kalayu")
	fmt.Println("After deleteing kalayu", userScores)

	//Read check if key exists
	value, exists := userScores["kalayu"]
	if exists {
		fmt.Println("kalayu's score is ", value)
	} else {
		fmt.Println("Error geting kalayu result")
	}

	//creating map with initial capactiy 20
	m := make(map[string]int, 10)
	for i := 0; i < 20; i++ {
		key := fmt.Sprint("key%d", i)
		m[key] = i
	}
	fmt.Println("capacity of m is :", len(m))

	//creating Nested maps
	employees := make(map[string]map[string]string)
	employees["emp1"] = map[string]string{
		"name":     "Kalayu",
		"Position": "Manager",
	}
	employees["emp2"] = map[string]string{
		"name":     "Hailom",
		"position": "developer",
	}

	fmt.Println("Employees:", employees)
	fmt.Println("emp1 Name", employees["emp1"]["name"])

	//Adding another entry to nested map
	employees["emp1"]["department"] = "engineering"
	fmt.Println("emp1 department", employees["emp1"]["department"])

	//map as function paramters
}

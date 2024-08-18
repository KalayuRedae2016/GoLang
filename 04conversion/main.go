package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Wellcome to our piza app")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate us our piza")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thank you for rating us", input)

	numrating, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("Adding Rating 1.0", numrating+1.0)
	}

}

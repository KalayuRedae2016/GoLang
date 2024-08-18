package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Wellcome to user Input"
	fmt.Println("Wellcome message:", welcome)

	reader := bufio.NewReader((os.Stdin))
	fmt.Println("Rating our piza product")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thank you for rating us", input)
}

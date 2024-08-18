package main

import "fmt"

type User struct { //define struct
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("wellcome to structs in golang")
	kalayu := User{"Kalayu", "kalayuredae2@gmail.com", true, 23}
	fmt.Println(kalayu)
	fmt.Printf("Kalayu Details:%+v\n", kalayu)

	fmt.Println(kalayu.Name)

}

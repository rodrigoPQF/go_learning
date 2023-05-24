package main

import "fmt"

type mytype int

var x mytype

func main() {
	fmt.Printf("Value = %v \nType = %T \n", x, x)

	x = 42

	fmt.Printf("Value = %v", x)

}

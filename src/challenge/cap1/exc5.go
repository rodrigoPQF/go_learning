package main

import "fmt"

type mytype int

var x mytype

var y int

func main() {
	fmt.Printf("Value X = %v \nType X = %T \n", x, x)
	x = 42
	fmt.Printf("Value X = %v \n", x)

	y = int(x)
	fmt.Printf("Value Y = %v \nType Y = %T \n", y, y)

}

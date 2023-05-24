package main

import "fmt"

func main() {
	x := 10
	fmt.Printf("%b , %d , %#x \n", x, x, x)

	y := x << 1
	fmt.Printf("%b , %d , %#x", y, y, y)

}

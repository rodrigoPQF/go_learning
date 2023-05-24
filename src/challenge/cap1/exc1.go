package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Print("X:", x, "\nY:", y, "\nZ:", z, "\n")

	fmt.Printf("X - type %T - value %v \n", x, x)
	fmt.Printf("Y - type %T - value %v \n", y, y)
	fmt.Printf("Z - type %T - value %v \n", z, z)

	fmt.Println("X - ", x)
	fmt.Println("Y - ", y)
	fmt.Println("Z - ", z)

}

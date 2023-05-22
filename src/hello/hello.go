package main

import "fmt"

var y = "strings" // Disponivel em todo o package - não da para usar ":="

func main() {
	x := 16 // Disponivel apenas nessa função

	z := true

	_, erros := fmt.Println("Hello World", "Oi galera", 100)

	fmt.Println(erros)                // Retorno de { nil }
	fmt.Printf("x: %v - %T \n", x, x) // $v = value - %T = type
	fmt.Printf("y: %v - %T \n", y, y)
	fmt.Printf("z: %v - %T \n", z, z)

	x, w := 20, 30

	fmt.Printf("x: %v - %T \n", x, x)
	fmt.Printf("x: %v - y: %v", x, w)

}

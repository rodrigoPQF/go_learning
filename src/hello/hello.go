package main

import "fmt"

var y = "strings" // Disponivel em todo o package - não da para usar ":="

func main() {
	z := 20
	anyThing(z)

}

func anyThing(x int) {
	// - Usando a variavel [y] declarado fora do escopo
	fmt.Println(y)
	// - Usando a variavel [x] dentro da função main ( z := 20)
	fmt.Println(x)
}

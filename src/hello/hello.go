package main

import "fmt"

var y = "strings" // Disponivel em todo o package - não da para usar ":="
var x int = 10
var k int

func main() {
	k = 10 // Só pode ser atribuido dentro de uma função
	z := 20
	// x = 'Test' - Go é staticamente tipado , não muda é possivel alterar o tipo
	anyThing(z)

}

func anyThing(x int) {
	// - Usando a variavel [y] declarado fora do escopo
	fmt.Println(y)
	// - Usando a variavel [x] dentro da função main ( z := 20)
	fmt.Println(x)
}

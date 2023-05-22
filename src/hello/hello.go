package main

import "fmt"

var y = "strings" // Disponivel em todo o package - não da para usar ":="
var x int = 10
var k int

func main() {
	// zeros()
	// pacoteFMT()
	// myType()
	// convert()
	// booleanFunc()
	// numericFunc()
	stringFunc()
}

func anyThing(x int) {
	// - Usando a variavel [y] declarado fora do escopo
	fmt.Println(y)
	// - Usando a variavel [x] dentro da função main ( z := 20)
	fmt.Println(x)
}

// Variáveis, Valores e TIpos - Valor Zero

var a int
var b float64
var c string
var d bool

func zeros() {
	fmt.Printf("A: %v \n", a)
	fmt.Printf("B: %v \n", b)
	fmt.Printf("C: %v \n", c)
	fmt.Printf("D: %v", d)

}

// Pacote FMT

func pacoteFMT() {
	x := "Oi"
	y := ", Bom dia"
	fmt.Print("Oi", "bom\n")     // Retornar bytes sem espaçamento
	fmt.Println("Oi", "bom dia") // Retorna bytes com espaçamento

	z := fmt.Sprint(x, y)  // Retorna uma String sem espaçamento
	z = fmt.Sprintln(x, y) // Retorna uma String com espaçamento

	fmt.Println(z)
}

// Criar tipagems

type hotdog int

var p hotdog = 10

func myType() {
	numeroComp := 10
	fmt.Printf("%T", p)
	p = hotdog(numeroComp) // p = numeroComp da erro
}

// Converter e nao coerção

func convert() {
	a := int(p)
	fmt.Printf("%T - %v", a, a)
}

// Tipo Booleano

func booleanFunc() {
	var x bool
	fmt.Println(x)
	x = true

	x = (10 < 100)
	y := 10 == 100
	z := 100 > 10
	fmt.Print(x, y, z)
}

func numericFunc() {
	// var num1 uint
	// var num2 int
	// var num3 float32
	// var num4 complex64
	// var num5 byte
	// var num6 rune
	a := "e"
	b := "é"
	c := "色"

	fmt.Printf("%v, %v ,%v \n", a, b, c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)

	fmt.Printf("%v, %v ,%v", d, e, f)

}

func stringFunc() {
	s := `Hello, 
	playground`
	sb := []byte(s)
	fmt.Printf("%v \n%T", sb, sb)

	for _, v := range sb {
		fmt.Printf("%v - %T - %#U - %#X \n", v, v, v, v)
	}
}

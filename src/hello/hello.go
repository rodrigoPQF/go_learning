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
	// stringFunc()
	// numericSystem()
	// constants()
	// iotaType()
	// deslocBit()
	// loopfor()
	// nestedLoop()
	// customWhile()
	// forContinue()
	// conditionIf()
	conditionSwitch()
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

func numericSystem() {
	a := 100
	fmt.Printf("%d\t%b\t%#x", a, a, a) //format verb
}

const l = 10 // Ele so se torna algo quando usado
const (
	ha = 10
	hb = 20
	hz = 30
)

var s int
var h float32

func constants() {
	s = l
	h = l
	fmt.Println(s)
	fmt.Println(ha, hb, hz)
}

//Iota

const (
	hx = iota + 1*10
	kz
	ka
)
const (
	ioa = iota
	iob = iota
	ioc = iota
)

// Usado quando não liga para os valores da constantes, so que seja diferente
// Tipo enum

func iotaType() {
	fmt.Println(hx, kz, ka)
}

// Deslocamento de bits

func deslocBit() {
	x := 1
	y := x << 1
	z := x >> 1
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", z)

}

func loopfor() {
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}

}

func nestedLoop() {
	for horas := 0; horas <= 24; horas++ {
		fmt.Print("Horas:", horas, "\n")
		for min := 0; min < 60; min++ {
			fmt.Print("Minutos: ", min, "\n")
			for segundos := 0; segundos < 60; segundos++ {
				fmt.Print("Segundos: ", segundos, "\n")
			}
		}
	}
}

func customWhile() {
	x := 0
	for x < 10 {
		fmt.Println("Esse while")
		x++
	}
	//Loop infinito
	for {
		fmt.Println("Loop infinito")
		if x < 10 {
			x++
			fmt.Println("Xis")
		} else {
			fmt.Println("Xiss ")
			break
		}
	}
}

func forContinue() {
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)

	}
}

func conditionIf() {
	x := 10
	if x < 100 {
		fmt.Println("Hello")
	}
	if !(x < 100) {
		fmt.Println("Hello is false")
	} else if x < 10 {
		fmt.Println("Hello is true")
	} else {
		fmt.Println("Hello is ...")
	}
}

func conditionSwitch() {
	x := 5
	y := 10
	switch true {
	case x == 5, y > 10:
		fmt.Println("xis igual a cinco")
		fallthrough
	case x == 3, y < 3:
		fmt.Println("XIS IGUAL A TRES")
	case x > 5:
		fmt.Println("xis e maior que cinco")
	default:
		fmt.Println("ta vazio")
	}

}

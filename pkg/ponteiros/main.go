package main

import "fmt"

type Ponteiro struct {
	Numero int32
}


func (p *Ponteiro) alterarPonteiro(num int32) {
	p.Numero = num
}

func main(){
	pont := Ponteiro{
		Numero: 3,
	}

	fmt.Println(pont.Numero)
	// out : 3
	pont.alterarPonteiro(4)
	
	fmt.Println(pont.Numero)
	// out : 4

	a := 10
	// Variavel a é apontado para um local da memoria e essa memoria tem um endereço

	fmt.Println(&a)
	// out: 0xc00004c768 - Esse é o endereço de memoria 0xc00004c768(10)

	var ponteiro *int = &a

	fmt.Println(ponteiro)
		// out: 0xc00004c768 - Var ponteiro aponta para o mesmo endereço 0xc00004c768(10)
	fmt.Println(*ponteiro)
	  // out: 10 - Exibe o valor que está armazenado

	*ponteiro = 50
	fmt.Println(a)
	fmt.Println(*ponteiro)
	// out: 50 - Altera o valor da memoria

	b := &a
	fmt.Println(b)
	// out: 50 
	*b = 60
	fmt.Println(b)
	fmt.Println(*ponteiro)
	fmt.Println(a)
	// out: 60 - Alterou todos os valores 

	variavel := 10
	abc(&variavel)
	fmt.Println(variavel)
	// out: 200 
}

func abc(a *int) {
	*a = 200
}
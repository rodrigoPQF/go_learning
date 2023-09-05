package ifElse

import "fmt"

func IfAndElse() {

	// Não utiliza parenteses 🦜
	if 2 > 1 {
		fmt.Println(true)
	} else if 5 > 2 {
		fmt.Println(false)
	} else {
		fmt.Println(false)
	}

	// If com init 💁 - A variavel nao é acessivel no codigo depois do If
	if retorno := RetornoString(); retorno == "Retorno" {
		fmt.Println(false)
	}
}

func RetornoString() string {
	return "Retorno"
}

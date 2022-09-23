package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Rodrigo"
	var idade = 21
	var versao float32 = 1.1
	fmt.Println("Olá mundo, sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variavel nome é", reflect.TypeOf((nome)))

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O endereço é", &comando)

	fmt.Println("O Comando escolhido foi", comando)

}

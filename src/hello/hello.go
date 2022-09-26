package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {

	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	// ############## IF NORMAL ##############
	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo Logs...")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo do Programa...")
	// } else {
	// 	fmt.Println("Não conheço este comando")
	// }

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do Programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)

	}

}

func exibeIntroducao() {
	nome := "Rodrigo"
	var idade = 21
	var versao float32 = 1.1
	fmt.Println("Olá mundo, sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("O tipo da variavel nome é", reflect.TypeOf((nome)))

}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O endereço é", &comandoLido)
	fmt.Println("O Comando escolhido foi", comandoLido)
	return comandoLido
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func saiDoPrograma() {
	fmt.Println("Saindo do Programa...")
}

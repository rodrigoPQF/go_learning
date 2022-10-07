package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	for {
		exibeNomes()
		// exibeIntroducao()
		// exibeMenu()

		_, idade := devolveNome()
		fmt.Println(idade)
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
			iniciaMonitoramento()
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

}

func devolveNome() (string, int) {
	nome := "Rodrigo"
	idade := 21
	return nome, idade

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

func iniciaMonitoramento() {
	sites := []string{"https://google.com", "https://www.alura.com.br", "https://www.alura.com.br"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Estou passando na posicao", i, "do meu slice e essa posicao tem o site:", site)
			testeSite(site)
		}
		time.Sleep(delay * time.Second)
	}

}
func testeSite(site string) {
	respon, _ := http.Get(site)

	if respon.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", respon.StatusCode)
	}
}
func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Bernardo"}
	nomes = append(nomes, "Aparecido")
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("Meu array tem ", len(nomes))
	fmt.Println("Meu slice tem capacidade para", cap(nomes), "itens")

}

package main

import "fmt"

// - Constantes
const meuNome = "Rodrigo"
const mundo = "Mundo"
const prefix_pt = "Olá, "
const prefix_en = "Hello, "
const prefix_es = "Hola, "
const prefix_fr = "Bonjour, "

// - Separando o "dominio" do resto - no caso o retorno da string
func Ola(nome, idioma string) string {
	if nome == "" {
		nome = mundo
	}
	return prefixSaudacao(idioma) + nome
}

func prefixSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case "portugues":
		prefixo = prefix_pt
	case "ingles":
		prefixo = prefix_en
	case "espanhol":
		prefixo = prefix_es
	case "frances":
		prefixo = prefix_fr
	default:
		prefixo = prefix_pt
	}
	return
}

func main() {
	fmt.Println(Ola(meuNome, "portugues"))
}

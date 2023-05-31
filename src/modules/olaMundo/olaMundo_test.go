package main

import "testing"

func TestOla(t *testing.T) {

	verificarMensagem := func(t *testing.T, resultado, esperando string) {
		// Usado para exibir o numero da linha caso o teste falha
		t.Helper()
		if resultado != esperando {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperando)
		}

	}
	t.Run("deverar dizer olá para min", func(t *testing.T) {
		resultado := Ola(meuNome, "ingles")
		esperando := "Hello, Rodrigo"
		verificarMensagem(t, resultado, esperando)

	})

	t.Run("deverar dizer olá para o mundo", func(t *testing.T) {
		resultado := Ola(mundo, "portugues")
		esperando := "Olá, Mundo"
		verificarMensagem(t, resultado, esperando)

	})

	t.Run("deverar dizer o idioma da saudação", func(t *testing.T) {
		resultado := Ola("Eloisa", "espanhol")
		esperando := "Hola, Eloisa"
		verificarMensagem(t, resultado, esperando)

	})
}

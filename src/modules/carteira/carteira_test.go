package carteira

import (
	"testing"
)

func TestCarteira(t *testing.T) {
	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))

	})
	t.Run("Retirar com saldo suficiente", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}
		erro := carteira.Retirar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))
		confirmarErrorInexistente(t, erro)
	})
	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		erro := carteira.Retirar(Bitcoin(100))

		confirmaSaldo(t, carteira, saldoInicial)
		confirmaError(t, erro, ErroSaldoInsuficiente)
	})

}
func confirmaError(t *testing.T, resultado error, esperado error) {
	t.Helper()
	if resultado == nil {
		t.Error("esperava um erro, mas nenhum ocorreu")
	}
	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

func confirmaSaldo(t *testing.T, carteira Carteira, esperado Bitcoin) {
	t.Helper()
	resultado := carteira.Saldo()
	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

func confirmarErrorInexistente(t *testing.T, resultado error) {
	t.Helper()
	if resultado != nil {
		t.Fatal("erro inesperado recebido")
	}
}

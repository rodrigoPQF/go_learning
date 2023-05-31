package soma

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {
	t.Run("coleção de slices qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3}
		resultado := Soma(numeros)
		esperado := 6

		if resultado != esperado {
			t.Errorf("resultado '%d', esperado '%d', dado '%v'", resultado, esperado, numeros)
		}
	})
}
func TestSomaTodoOResto(t *testing.T) {

	verificarSomas := func(t *testing.T, resultado, esperado []int) {
		t.Helper()
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado '%v', esperado '%d'", resultado, esperado)
		}

	}
	t.Run("faz as somas de alguns slices", func(t *testing.T) {

		resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}

		verificarSomas(t, resultado, esperado)
	})
	t.Run("faz as somas de slices vazios de forma segura", func(t *testing.T) {

		resultado := SomaTodoOResto([]int{}, []int{0, 9})
		esperado := []int{0, 9}

		verificarSomas(t, resultado, esperado)

	})
}

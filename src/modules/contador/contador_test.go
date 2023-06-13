package contador

import (
	"sync"
	"testing"
)

func TestContador(t *testing.T) {
	t.Run("incrementar o contador 3 vezes resulta no valor 3", func(t *testing.T) {
		contador := Contador{}
		contador.Incremeta()
		contador.Incremeta()
		contador.Incremeta()

		verificaContador(t, &contador, 3)
	})

	t.Run("roda concorrentemente em segurança", func(t *testing.T) {
		contagemEsperada := 1000
		contador := Contador{}

		var wg sync.WaitGroup
		wg.Add(contagemEsperada)

		for i := 0; i < contagemEsperada; i++ {
			go func(w *sync.WaitGroup) {
				contador.Incremeta()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		verificaContador(t, &contador, contagemEsperada)
	})
}

func verificaContador(t *testing.T, resultado *Contador, esperado int) {
	t.Helper()
	if resultado.Valor() != esperado {
		t.Errorf("resultado %d, esperado %d", resultado.Valor(), esperado)
	}
}

package corredor

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCorredor(t *testing.T) {
	t.Run("compara a velocidade de servidores, retornando o endereço do mais rápido", func(t *testing.T) {
		servidorLento := criarServidorComAtrasado(20 * time.Millisecond)
		servidorRapido := criarServidorComAtrasado(0 * time.Millisecond)

		defer servidorLento.Close()
		defer servidorRapido.Close()

		URLLenta := servidorLento.URL
		URLRapida := servidorRapido.URL

		esperado := URLRapida
		resultado, err := Corredor(URLLenta, URLRapida)

		if err != nil {
			t.Fatalf("não esperava um erro, mas obteve um %v", err)
		}

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})
	t.Run("retornar um erro se o servidor nao responder dentro de 10s", func(t *testing.T) {
		servidorA := criarServidorComAtrasado(11 * time.Millisecond)
		servidorB := criarServidorComAtrasado(12 * time.Millisecond)

		defer servidorA.Close()
		defer servidorB.Close()

		_, err := Corredor(servidorA.URL, servidorB.URL)

		if err == nil {
			t.Error("esperava um erro, mas não obtive um")
		}

	})

}

func criarServidorComAtrasado(atraso time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(atraso)
		w.WriteHeader(http.StatusOK)
	}))
}

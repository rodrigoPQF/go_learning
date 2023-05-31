package calculos

import "testing"

func TestPerimetro(t *testing.T) {
	retanguloSize := Retangulo{10.00, 10.00}
	resultado := Perimetro(retanguloSize)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("resultado %.2f esperado %.2f", resultado, esperado) // .2f para float64 e duas casas decimais
	}
}

func TestArea(t *testing.T) {
	verificaArea := func(t *testing.T, forma Forma, esperado float64) {
		t.Helper()
		resultado := forma.Area()

		if resultado != esperado {
			t.Errorf("resultado %.2f esperado %.2f", resultado, esperado) // .2f para float64 e duas casas decimais
		}

	}
	t.Run("retangulos", func(t *testing.T) {
		retanguloSize := Retangulo{10.00, 10.00}
		verificaArea(t, retanguloSize, 100.00)

	})

	t.Run("circulo", func(t *testing.T) {
		circuloSize := Circulo{10.00}
		verificaArea(t, circuloSize, 314.1592653589793)

	})

}

func TestTableArea(t *testing.T) {
	testesArea := []struct {
		forma    Forma
		esperado float64
	}{
		{forma: Retangulo{10, 6}, esperado: 60.0},
		{forma: Circulo{10}, esperado: 314.1592653589793},
		{forma: Triangulo{12, 6}, esperado: 36.0},
	}

	for _, tt := range testesArea {
		resultado := tt.forma.Area()

		if resultado != tt.esperado {
			t.Errorf("resultado %.2f esperado %.2f", resultado, tt.esperado) // .2f para float64 e duas casas decimais
		}
	}

}

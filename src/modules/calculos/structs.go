package calculos

type Retangulo struct {
	Largura float64
	Altura  float64
}

type Circulo struct {
	Raio float64
}

type Triangulo struct {
	Base   float64
	Altura float64
}

// Retangulo tem um método chamado Area que retorna um float64, então satisfaz a interface Forma.
// Circulo tem um método chamado Area que retorna um float64, então satisfaz a interface Forma.
// string não tem esse método, então não satisfaz a interface.
// etc.

type Forma interface {
	Area() float64
}

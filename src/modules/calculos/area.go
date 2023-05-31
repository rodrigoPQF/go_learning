package calculos

import "math"

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}

/* Mesma coisa que :
 func () {
	area(){

	}
}
*/

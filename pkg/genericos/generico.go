package genericos

import "fmt"

// Módulo de aprendizado de genericos do go

// Struct generica
type Listas[T any] struct {
	Nome      T
	Sobrenome T
}

// Token
type Token int64

func Generico() {
	// Só aceitara o tipo passado T = string nesse caso
	listas := Listas[string]{
		Nome:      "Rodrigo",
		Sobrenome: "Pereira",
	}

	fmt.Println(listas)

	var token Token = 32
	Numero(token, token)

	CompararNumeros(34, 43.2)

	m := map[int]string{1: "2", 3: "4", 5: "6"}
	r := CompararTipos(m)
	fmt.Println(r)
}

// Token ~ e necessario para implementar tipos que tem a mesma implementação
func Numero[T ~int64](arg1 T, arg2 T) {
	fmt.Println(arg1, arg2)
}

func CompararTipos[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type Numeros interface {
	int64 | float64 | int32 | float32
}

func CompararNumeros[T Numeros](arg1 T, arg2 T) {
	if arg1 > arg2 {
		fmt.Println("Maior")
	}
}

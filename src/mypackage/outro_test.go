package mypackage_test

import (
	"module/src/mypackage"
	"testing"
)

type testes struct {
	data   []int
	answer int
}

func TestSomaEmTabela(t *testing.T) {
	tests := []testes{
		{data: []int{1, 2, 3}, answer: 6},
		{[]int{10, 11, 12}, 33},
		{[]int{-5, 0, 5, 10}, 10},
	}
	for _, v := range tests {
		z := mypackage.Somas(v.data...)
		if z != v.answer {
			t.Error("Expected: ", v.answer, "Got: ", z)
		}
	}
}

func TestSoma(t *testing.T) {
	teste := mypackage.Somas(3, 2, 1)
	resultado := 6
	if teste != 6 {
		t.Error("Expected: ", resultado, "Got: ", teste)
	}
}

func TestSoma2(t *testing.T) {
	teste := mypackage.Somas(3, 2, 1)
	resultado := 6
	if teste != 6 {
		t.Error("Expected: ", resultado, "Got: ", teste)
	}
}

func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mypackage.Soma(1, 1)
	}
}

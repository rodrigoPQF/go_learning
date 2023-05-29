package mypackage

import "fmt"

func Printando(a ...interface{}) {
	fmt.Println("Essa func e main")
}

func Soma(x int, y int) int {
	return x + y
}

func Somas(i ...int) int {
	total := 0
	for _, v := range i {
		total = total + v
	}
	return total
}

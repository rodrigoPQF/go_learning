package main

import "fmt"

func main() {
	const (
		_ = 1994 + iota
		a
		b
		c
	)
	fmt.Println(a, b, c)
}

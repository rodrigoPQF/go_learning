package main

import (
	"fmt"
	"io"
	"net/http"
)

func Cumprimenta(escritor io.Writer, nome string) {
	fmt.Fprintf(escritor, "Olá, %s", nome)
}

func HandlerMeuCumprimento(w http.ResponseWriter, r *http.Request) {
	Cumprimenta(w, "mundo")
}

func main() {
	err := http.ListenAndServe(":3366", http.HandlerFunc(HandlerMeuCumprimento))
	if err != nil {
		fmt.Println(err)
	}
}

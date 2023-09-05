package variaveis

import "fmt"

func Variaveis() {

	var inteiros32 int32
	var inteiros64 int64

	var flutuantes32 float32
	var flutuantes64 float64

	var literais string

	var array map[string]string

	fmt.Println(inteiros32, inteiros64, flutuantes32, flutuantes64, literais, array)

	// Auto tipagem do GO 😺
	issoEString := "Olá Mundo"
	issoEFloat := 92.2
	issoEInt := 2
	issoEArray := map[string]string{
		"Olá": "Mundo",
	}

	fmt.Printf("%T", issoEString)
	fmt.Printf("%T", issoEFloat)
	fmt.Printf("%T", issoEInt)
	fmt.Printf("%T", issoEArray)

	issoPrivado()
	IssoPublico()

	// Tipo Any do Go
	var issoEumAny interface{}

	fmt.Println(issoEumAny)

	// Arrays tem tamanho definido
	var arrays [4]string = [4]string{"1", "2", "3", "4"}
	fmt.Println(arrays, "Tamanho", len(arrays), "Capacidade", cap(arrays))

	// Slices não tem tamanho definido
	var slices []string = []string{"1", "2", "3", "4", "5", "6"}
	fmt.Println(arrays, "Tamanho", len(slices), "Capacidade", cap(slices))
	slices = append(slices, "+1")
	fmt.Println(arrays, "Tamanho", len(slices), "Capacidade", cap(slices))

}

// Função/Variaveis privadas é so colocar em smallcase 🔸

// Esse é privado
func issoPrivado() {}

// Esse é publico
func IssoPublico() {}

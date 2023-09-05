package switchCases

import "fmt"

func SwitchCases() {

	test := "Teste"

	// Switch cases 🦢
	switch test {
	case "test1":
		fmt.Println("Caiu na primeira")
		// Isso aqui ele vai cair na segunda independente da condição
		fallthrough
	case "test2", "test3":
		fmt.Println("Caiu na segunda")
	case "Teste":
		fmt.Println("Caiu na CERTA")
	default:
		fmt.Println("Xiii caiu em nada")
	}

}

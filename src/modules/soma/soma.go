package soma

func Soma(numeros []int) int {
	soma := 0

	for _, v := range numeros {
		soma += v
	}
	return soma

}

func SomaTodoOResto(numerosSoma ...[]int) []int {
	var somas []int

	for _, numeros := range numerosSoma {
		if len(numeros) == 0 {
			somas = append(somas, 0)
		} else {
			final := numeros[1:]
			somas = append(somas, Soma(final))
		}
	}
	return somas
}

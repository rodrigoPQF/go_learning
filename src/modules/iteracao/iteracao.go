package iteracao

const quantidadeRepeticoes = 5

func Repetir(caracteres string) string {
	var repeticoes string
	for i := 0; i < quantidadeRepeticoes; i++ {
		repeticoes += caracteres

	}
	return repeticoes

}

package contagem

import (
	"fmt"
	"io"
	"os"
	"time"
)

const ultimaPalavra = "Go!"
const inicioContagem = 3
const escrita = "escrita"
const pausa = "pausa"

func (s *SleeperSpy) Sleep() {
	s.Chamadas++
}

func (d *SleeperPadrao) Sleep() {
	time.Sleep(1 * time.Second)
}
func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		fmt.Fprintln(saida, i)
	}
	for i := inicioContagem; i > 0; i-- {
		fmt.Fprintln(saida, i)
	}

	fmt.Fprint(saida, ultimaPalavra)
}

func (s *SpyContagemOperacoes) Pausa() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

func (t *TempoSpy) Pausa(duracao time.Duration) {
	t.duracaoPausa = duracao
}

func (s *SleeperConfiguravel) Pausa() {
	s.pausa(s.duracao)
}

func main() {
	sleeper := &SleeperPadrao{}
	Contagem(os.Stdout, sleeper)
}

package contagem

import "time"

type Sleeper interface {
}

type SleeperSpy struct {
	Chamadas int
}

type SpyContagemOperacoes struct {
	Chamadas []string
}

type SleeperPadrao struct{}

type SleeperConfiguravel struct {
	duracao time.Duration
	pausa   func(time.Duration)
}

type TempoSpy struct {
	duracaoPausa time.Duration
}

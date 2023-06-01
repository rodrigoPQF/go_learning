package carteira

type Bitcoin int

type Carteira struct {
	saldo Bitcoin
}

type Stringer interface {
	Stringer() string
}

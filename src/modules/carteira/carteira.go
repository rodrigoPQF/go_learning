package carteira

import (
	"errors"
	"fmt"
)

var ErroSaldoInsuficiente = errors.New("não é possivel retirar: saldo insuficiente")

func (c *Carteira) Depositar(quantidade Bitcoin) {

	c.saldo += quantidade
}

func (c *Carteira) Retirar(quantidade Bitcoin) error {

	if quantidade > c.saldo {
		return ErroSaldoInsuficiente
	}

	c.saldo -= quantidade
	return nil
}
func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

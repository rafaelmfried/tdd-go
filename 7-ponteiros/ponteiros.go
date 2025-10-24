package ponteiros

import (
	"fmt"
)

var ErroSaldoInsuficiente = fmt.Errorf("valor maior que saldo disponivel")

type Bitcoin int
type Carteira struct {
	saldo int
}

func NewCarteira(saldo int) *Carteira {
	return &Carteira{saldo: saldo}
}

func (c *Carteira) Depositar(valor int) {
	c.saldo += valor
}

func (c *Carteira) Saldo() int {
	return c.saldo
}

func (c *Carteira) Retirar(valor int) error {
	if valor > c.saldo {
		return ErroSaldoInsuficiente
	}

	c.saldo -= valor
	return nil
}

type CarteiraBitcoin struct {
	saldo Bitcoin
}

func NewBitcoinCarteira(saldo Bitcoin) *CarteiraBitcoin {
	return &CarteiraBitcoin{saldo: saldo}
}

func (c *CarteiraBitcoin) Depositar(valor Bitcoin) {
	c.saldo += valor
}

func (c *CarteiraBitcoin) Saldo() Bitcoin {
	return c.saldo
}

func (c *CarteiraBitcoin) Retirar(valor Bitcoin) error {
	if valor > c.saldo {
		return ErroSaldoInsuficiente
	}

	c.saldo -= valor
	return nil
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
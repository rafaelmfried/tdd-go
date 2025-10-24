package sync

func NovoContador() *Contador {
	return &Contador{}
}

type Contador struct {
	valor int
}

func (c *Contador) Incrementar() {
	c.valor++
}

func (c *Contador) Valor() int {
	return c.valor
}
/*
go vet para ajudar a revelar problemas comuns em codigo concorrente
*/
package sync

import "sync"

func NovoContador() *Contador {
	return &Contador{}
}

type Contador struct {
	sync.Mutex
	valor int
}

func (c *Contador) Incrementar() {
	c.Lock()
	defer c.Unlock()
	c.valor++
}

func (c *Contador) Valor() int {
	c.Lock()
	defer c.Unlock()
	return c.valor
}
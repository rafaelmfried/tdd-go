package cli

import (
	"bufio"
	"io"
	"strings"
	"tdd/application/1-http/linha-de-comando/armazenamento"
)

type CLI struct {
	armazenamento armazenamento.ArmazenamentoJogador
	in io.Reader
}

func NovoCLI(armazenamento armazenamento.ArmazenamentoJogador, in io.Reader) *CLI {
	return &CLI{
		armazenamento: armazenamento,
		in:           in,
	}
}

func (c *CLI) JogarPoquer() {
	reader := bufio.NewScanner(c.in)
	reader.Scan()
	vencedor := extrairVencedor(reader.Text())
	c.armazenamento.RegistrarVitoria(vencedor)
}

func extrairVencedor(input string) string {
	return strings.Replace(input, " venceu", "", 1)
}
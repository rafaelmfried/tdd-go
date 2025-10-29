package cli

import (
	"bufio"
	"io"
	"strings"
	"tdd/application/1-http/linha-de-comando/armazenamento"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

type CLI struct {
	armazenamento armazenamento.ArmazenamentoJogador
	in io.Reader
	blindAlerter BlindAlerter
}

func NovoCLI(armazenamento armazenamento.ArmazenamentoJogador, in io.Reader, blindAlerter BlindAlerter) *CLI {
	return &CLI{
		armazenamento: armazenamento,
		in:           in,
		blindAlerter: blindAlerter,
	}
}

func (c *CLI) JogarPoquer() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		c.blindAlerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + 10 * time.Minute
	}

	input := readLine(c.in)
	vencedor := extrairVencedor(input)
	c.armazenamento.RegistrarVitoria(vencedor)
}

func extrairVencedor(input string) string {
	return strings.Replace(input, " venceu", "", 1)
}

func readLine(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	return scanner.Text()
}
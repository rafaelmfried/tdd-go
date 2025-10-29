package cli

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"tdd/application/1-http/linha-de-comando/armazenamento"
	"tdd/application/1-http/linha-de-comando/poker"
	"time"
)

const PlayerPrompt = "Please enter the number of players: "

type CLI struct {
	armazenamento armazenamento.ArmazenamentoJogador
	in io.Reader
	out io.Writer
	blindAlerter poker.BlindAlerter
}

func NovoCLI(armazenamento armazenamento.ArmazenamentoJogador, in io.Reader, out io.Writer, blindAlerter poker.BlindAlerter) *CLI {
	return &CLI{
		armazenamento: armazenamento,
		in:           in,
		out:          out,
		blindAlerter: blindAlerter,
	}
}

func (c *CLI) JogarPoquer() {
	fmt.Fprint(c.out, PlayerPrompt)
	c.scheduleBlindAlerts(5)
	input := readLine(c.in)
	vencedor := extrairVencedor(input)
	c.armazenamento.RegistrarVitoria(vencedor)
}

func (c *CLI) scheduleBlindAlerts(numPlayers int) {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	interval := time.Duration(5 + numPlayers) * time.Minute
	for _, blind := range blinds {
		c.blindAlerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + interval
	}
}

func extrairVencedor(input string) string {
	return strings.Replace(input, " venceu", "", 1)
}

func readLine(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	return scanner.Text()
}
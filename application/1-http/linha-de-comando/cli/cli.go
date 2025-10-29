package cli

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"tdd/application/1-http/linha-de-comando/poker"
)

const PlayerPrompt = "Please enter the number of players: "

type CLI struct {
	in *bufio.Reader
	out io.Writer
	game poker.Game
}

func NovoCLI(in io.Reader, out io.Writer, game poker.Game) *CLI {
	return &CLI{
		in:           bufio.NewReader(in),
		out:          out,
		game:        game,
	}
}

func (c *CLI) JogarPoquer() {
	fmt.Fprint(c.out, PlayerPrompt)

	numberOfPlayersInput := c.readLine()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))
	if err != nil {
		fmt.Fprintf(c.out, "entrada invalida para numero de jogadores '%s'\n", numberOfPlayersInput)
		return
	}
	
	c.game.Start(numberOfPlayers)
	
	winnerInput := c.readLine()
	vencedor := extrairVencedor(winnerInput)

	c.game.Finish(vencedor)
}

func extrairVencedor(input string) string {
	return strings.Replace(input, " venceu", "", 1)
}

func (c *CLI) readLine() string {
	scanner := bufio.NewScanner(c.in)
	scanner.Scan()
	return scanner.Text()
}
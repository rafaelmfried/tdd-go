package cli

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"tdd/application/1-http/linha-de-comando/poker"
)

const PlayerPrompt = "Please enter the number of players: "

const BadPlayerInputErrMsg = "bad value received for number of players, please try again with a number"
var ErrBadPlayerInput = errors.New(BadPlayerInputErrMsg)

const BadWinnerInputErrMsg = "bad value received for winner, please try again with format 'Name venceu'"
var ErrBadWinnerInput = errors.New(BadWinnerInputErrMsg)

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
	numberOfPlayers, err := strconv.Atoi(numberOfPlayersInput)
	if err != nil {
		fmt.Fprint(c.out, ErrBadPlayerInput)
		return
	}
	
	c.game.Start(numberOfPlayers)
	
	winnerInput := c.readLine()
	vencedor, err := extrairVencedor(winnerInput)
	if err != nil {
		fmt.Fprint(c.out, err)
		return
	}

	c.game.Finish(vencedor)
}

func extrairVencedor(input string) (string, error) {
	if !validatedWinnerInput(input) {
		return "", ErrBadWinnerInput
	}
	return strings.Replace(input, " venceu", "", 1), nil
}

func validatedWinnerInput(input string) bool {
	return strings.Contains(input, " venceu")
}

func (c *CLI) readLine() string {
    reader := bufio.NewReader(c.in)
    line, _ := reader.ReadString('\n')
    return strings.TrimSpace(line)
}
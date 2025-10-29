/*
        assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg) OK
        assertGameNotStarted(t, game) OK
        assertFinishCalledWith(t, game, "Cleo") OK
        assertGameStartedWith(t, game, 8) OK
        in := userSends("3", "Chris wins")

				+ cases: What happens if instead of putting Ruth wins the user puts in Lloyd is a killer ?
*/

package cli_test

import (
	"bytes"
	"strings"
	"tdd/application/1-http/linha-de-comando/cli"
	"tdd/application/1-http/linha-de-comando/helpers"
	"testing"
	"time"
)

const PlayerPrompt = cli.PlayerPrompt
const BadPlayerInputErrMsg = cli.BadPlayerInputErrMsg
const BadWinnerInputErrMsg = cli.BadWinnerInputErrMsg

// SpyBlindAlerter para testes
type SpyBlindAlerter struct{}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	// Implementação vazia para o teste
}

func TestCLI(t *testing.T) {
	t.Run("testa chamada de vitorias pela linha de comando", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := helpers.UserSends("5", "Chris venceu")
		game := helpers.NovoGameSpy()
		cli := cli.NovoCLI(in, stdout, game)

		cli.JogarPoquer()

		wantPrompt := PlayerPrompt
		
		assertMessagesSentToUser(t, stdout, wantPrompt)
		assertGameStartedWith(t, game, 5)
		verificaVitoriaJogador(t, game, "Chris")
	})

	t.Run("recorda vencedor cleo digitado pelo usuario", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := helpers.UserSends("5", "Cleo venceu")
		game := helpers.NovoGameSpy()

		cli := cli.NovoCLI(in, stdout, game)
		cli.JogarPoquer()

		wantPrompt := PlayerPrompt
		
		assertMessagesSentToUser(t, stdout, wantPrompt)
		assertGameStartedWith(t, game, 5)
		assertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := helpers.UserSends("7", "Ruth venceu")
		game := helpers.NovoGameSpy()

		cli := cli.NovoCLI(in, stdout, game)
		cli.JogarPoquer()

		wantPrompt := PlayerPrompt
		assertMessagesSentToUser(t, stdout, wantPrompt)
		assertGameStartedWith(t, game, 7)
	})

	t.Run("deve retornar um um erro caso usuario coloque um valor n valido para quantidade de jogadores", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := helpers.UserSends("Pies")
		game := helpers.NovoGameSpy()

		cli := cli.NovoCLI(in, stdout, game)

		cli.JogarPoquer()

		assertGameNotStarted(t, game)
	})

	t.Run("deve retornar uma mensagem de erro caso usuario coloque um valor n valido para quantidade de jogadores", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := helpers.UserSends("Pies", "Ruth venceu")
		game := helpers.NovoGameSpy()

		cli := cli.NovoCLI(in, stdout, game)

		cli.JogarPoquer()

		wantPrompt := PlayerPrompt + BadPlayerInputErrMsg

		assertMessagesSentToUser(t, stdout, wantPrompt)
		assertGameNotStarted(t, game)
	})

	t.Run("deve retornar uma mensagem de erro caso o usuario n coloque o sufixo venceu depois do nome do vencedor", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := helpers.UserSends("5", "Cleo killer")
		game := helpers.NovoGameSpy()

		cli := cli.NovoCLI(in, stdout, game)

		cli.JogarPoquer()

		wantPrompt := PlayerPrompt + BadWinnerInputErrMsg

		assertMessagesSentToUser(t, stdout, wantPrompt)
	})
}

func verificaVitoriaJogador(t *testing.T, game *helpers.GameSpy, vencedor string) {
    t.Helper()

		if game.FinishedWith != vencedor {
				t.Errorf("nao armazenou o vencedor correto, recebi '%s' esperava '%s'", game.FinishedWith, vencedor)
		}
}

func assertMessagesSentToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
    t.Helper()
    want := strings.Join(messages, "")
    got := stdout.String()
    if got != want {
        t.Errorf("got '%s' sent to stdout but expected %+v", got, messages)
    }
}

func assertGameNotStarted(t *testing.T, game *helpers.GameSpy) {
		t.Helper()
		if game.StartCalled {
			t.Errorf("expected game not to have started")
		}
}

func assertGameStartedWith(t *testing.T, game *helpers.GameSpy, expectedPlayers int) {
	t.Helper()
	if game.StartCalledWith != expectedPlayers {
		t.Errorf("expected game to have started with %d players but got %d", expectedPlayers, game.StartCalledWith)
	}
}

func assertFinishCalledWith(t *testing.T, game *helpers.GameSpy, expectedWinner string) {
	t.Helper()
	if game.FinishedWith != expectedWinner {
		t.Errorf("expected game to have finished with winner %q but got %q", expectedWinner, game.FinishedWith)
	}
}
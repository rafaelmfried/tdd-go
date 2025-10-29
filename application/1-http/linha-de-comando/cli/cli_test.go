package cli_test

import (
	"bytes"
	"strings"
	"tdd/application/1-http/linha-de-comando/cli"
	"tdd/application/1-http/linha-de-comando/helpers"
	"testing"
	"time"
)

const PlayerPrompt = "Please enter the number of players: "
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"

// SpyBlindAlerter para testes
type SpyBlindAlerter struct{}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	// Implementação vazia para o teste
}

func TestCLI(t *testing.T) {
	t.Run("testa chamada de vitorias pela linha de comando", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("5\n")
		game := helpers.NovoGameSpy()
		cli := cli.NovoCLI(in, stdout, game)

		cli.JogarPoquer()

		game.Start(5)

		gotPrompt := stdout.String()
		wantPrompt := PlayerPrompt

		game.Finish("Chris")
		
		if gotPrompt != wantPrompt {
			t.Errorf("esperava prompt '%s', mas obteve '%s'", wantPrompt, gotPrompt)
		}

		if game.StartCalledWith != 5 {
			t.Errorf("esperava que o jogo fosse iniciado com 5 jogadores, mas foi iniciado com %d", game.StartCalledWith)
		}


		verificaVitoriaJogador(t, game, "Chris")
	})

	t.Run("recorda vencedor cleo digitado pelo usuario", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("5\n")
		game := helpers.NovoGameSpy()

		cli := cli.NovoCLI(in, stdout, game)
		cli.JogarPoquer()

		gotPrompt := stdout.String()
		wantPrompt := PlayerPrompt

		game.Finish("Cleo")
		
		if gotPrompt != wantPrompt {
			t.Errorf("esperava prompt '%s', mas obteve '%s'", wantPrompt, gotPrompt)
		}

		if game.StartCalledWith != 5 {
			t.Errorf("esperava que o jogo fosse iniciado com 5 jogadores, mas foi iniciado com %d", game.StartCalledWith)
		}
		
		verificaVitoriaJogador(t, game, "Cleo")
	})

	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := helpers.NovoGameSpy()

		cli := cli.NovoCLI(in, stdout, game)
		cli.JogarPoquer()

		gotPrompt := stdout.String()
		wantPrompt := PlayerPrompt

		if gotPrompt != wantPrompt {
				t.Errorf("got '%s', want '%s'", gotPrompt, wantPrompt)
		}

		if game.StartCalledWith != 7 {
				t.Errorf("wanted Start called with 7 but got %d", game.StartCalledWith)
		}
	})

	t.Run("deve retornar um um erro caso usuario coloque um valor n valido para quantidade de jogadores", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := helpers.NovoGameSpy()

		cli := cli.NovoCLI(in, stdout, game)

		cli.JogarPoquer()

		if game.StartCalled {
			t.Errorf("game should not have started")
		}
	})

	t.Run("deve retornar uma mensagem de erro caso usuario coloque um valor n valido para quantidade de jogadores", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("you're so silly\n")
		game := helpers.NovoGameSpy()

		cli := cli.NovoCLI(in, stdout, game)

		cli.JogarPoquer()

		gotPrompt := stdout.String()

		wantPrompt := PlayerPrompt + BadPlayerInputErrMsg

		if gotPrompt != wantPrompt {
				t.Errorf("got '%s', want '%s'", gotPrompt, wantPrompt)
		}
	})

}

func verificaVitoriaJogador(t *testing.T, game *helpers.GameSpy, vencedor string) {
    t.Helper()

    if game.FinishedWith != vencedor {
        t.Errorf("nao armazenou o vencedor correto, recebi '%s' esperava '%s'", game.FinishedWith, vencedor)
    }
}

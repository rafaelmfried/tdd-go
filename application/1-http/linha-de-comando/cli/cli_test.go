package cli_test

import (
	"bytes"
	"strings"
	"tdd/application/1-http/linha-de-comando/cli"
	"tdd/application/1-http/linha-de-comando/helpers"
	"testing"
	"time"
)

// SpyBlindAlerter para testes
type SpyBlindAlerter struct{}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	// Implementação vazia para o teste
}

func TestCLI(t *testing.T) {
	t.Run("testa chamada de vitorias pela linha de comando", func(t *testing.T) {
		in := strings.NewReader("Chris venceu\n")
		stdout := &bytes.Buffer{}
		armazenamentoJogado := &helpers.EsbocoArmazenamentoJogador{}
		dummyBlindAlerter := &SpyBlindAlerter{}
		cli := cli.NovoCLI(armazenamentoJogado, in, stdout, dummyBlindAlerter)

		cli.JogarPoquer()

		verificaVitoriaJogador(t, armazenamentoJogado, "Chris")
	})

	t.Run("recorda vencedor cleo digitado pelo usuario", func(t *testing.T) {
		in := strings.NewReader("Cleo venceu\n")
		stdout := &bytes.Buffer{}
		armazenamentoJogado := &helpers.EsbocoArmazenamentoJogador{}
		dummyBlindAlerter := &SpyBlindAlerter{}
		cli := cli.NovoCLI(armazenamentoJogado, in, stdout, dummyBlindAlerter)

		cli.JogarPoquer()

		verificaVitoriaJogador(t, armazenamentoJogado, "Cleo")
	})
}

func verificaVitoriaJogador(t *testing.T, armazenamento *helpers.EsbocoArmazenamentoJogador, vencedor string) {
    t.Helper()

    if len(armazenamento.RegistrosVitorias) != 1 {
        t.Fatalf("recebi %d chamadas de GravarVitoria esperava %d", len(armazenamento.RegistrosVitorias), 1)
    }

    if armazenamento.RegistrosVitorias[0] != vencedor {
        t.Errorf("nao armazenou o vencedor correto, recebi '%s' esperava '%s'", armazenamento.RegistrosVitorias[0], vencedor)
    }
}

package cli_test

import (
	"strings"
	"tdd/application/1-http/linha-de-comando/armazenamento"
	"tdd/application/1-http/linha-de-comando/cli"
	"tdd/application/1-http/linha-de-comando/liga"
	"testing"
)

// Depois achar alguma maneira de tirar isso daqui e colocar em uma maneira que possa usar nos testes

type EsbocoArmazenamentoJogador struct {
	pontuacoes map[string]int
	registrosVitorias []string
	liga liga.Liga
}

var ErrJogadorNotFound = armazenamento.ErrJogadorNotFound

func (e *EsbocoArmazenamentoJogador) ObterPontuacaoJogador(nome string) (pontuacao int, err error) {
	pontuacao = e.pontuacoes[nome]
	if pontuacao == 0 {
		return 0, ErrJogadorNotFound
	}
	return pontuacao, nil
}

func (e *EsbocoArmazenamentoJogador) RegistrarVitoria(nome string) {
	e.registrosVitorias = append(e.registrosVitorias, nome)
}

func (e *EsbocoArmazenamentoJogador) ObterLiga() liga.Liga  {
	return e.liga
}

func TestCLI(t *testing.T) {
	t.Run("testa chamada de vitorias pela linha de comando", func(t *testing.T) {
		in := strings.NewReader("Chris venceu\n")
		armazenamentoJogado := &EsbocoArmazenamentoJogador{}
		cli := cli.NovoCLI(armazenamentoJogado, in)

		cli.JogarPoquer()

		verificaVitoriaJogador(t, armazenamentoJogado, "Chris")
	})

	t.Run("recorda vencedor cleo digitado pelo usuario", func(t *testing.T) {
		in := strings.NewReader("Cleo venceu\n")
		armazenamentoJogado := &EsbocoArmazenamentoJogador{}
		cli := cli.NovoCLI(armazenamentoJogado, in)

		cli.JogarPoquer()

		verificaVitoriaJogador(t, armazenamentoJogado, "Cleo")
	})
}

func verificaVitoriaJogador(t *testing.T, armazenamento *EsbocoArmazenamentoJogador, vencedor string) {
    t.Helper()

    if len(armazenamento.registrosVitorias) != 1 {
        t.Fatalf("recebi %d chamadas de GravarVitoria esperava %d", len(armazenamento.registrosVitorias), 1)
    }

    if armazenamento.registrosVitorias[0] != vencedor {
        t.Errorf("nao armazenou o vencedor correto, recebi '%s' esperava '%s'", armazenamento.registrosVitorias[0], vencedor)
    }
}

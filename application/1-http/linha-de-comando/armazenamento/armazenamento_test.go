package armazenamento_test

import (
	. "tdd/application/1-http/linha-de-comando/armazenamento"
	helpers "tdd/application/1-http/linha-de-comando/helpers"
	"tdd/application/1-http/linha-de-comando/liga"
	"testing"
)

func TestSistemaDeArquivoDeArmazenamentoDoJogador(t *testing.T) {
	t.Run("liga eh carregada de um leitor", func(t *testing.T) {
		conteudoArquivo := `[
			{"Nome": "Rafael", "Pontos": 10},
			{"Nome": "Vanessa", "Pontos": 20},
			{"Nome": "Pedro", "Pontos": 30}
		]`

		bancoDeDados, limpaBancoDeDados := helpers.CriarArquivoTemporario(t, conteudoArquivo)
		defer limpaBancoDeDados()

		armazenamento, _ := NovoArmazenamentoJogadorDoArquivo(bancoDeDados)

		recebido := armazenamento.ObterLiga()

		esperado := []liga.Jogador{
			{Nome: "Pedro", Pontos: 30},
			{Nome: "Vanessa", Pontos: 20},
			{Nome: "Rafael", Pontos: 10},
		}

		helpers.VerificaLiga(t, recebido, esperado)
	})

	t.Run("pegar pontuacao do jogador do arquivo", func(t *testing.T) {
		conteudoArquivo := `[
			{"Nome": "Rafael", "Pontos": 10},
			{"Nome": "Vanessa", "Pontos": 20},
			{"Nome": "Pedro", "Pontos": 30}
		]`
		bancoDeDados, limpaBancoDeDados := helpers.CriarArquivoTemporario(t, conteudoArquivo)
		defer limpaBancoDeDados()

		armazenamento, _ := NovoArmazenamentoJogadorDoArquivo(bancoDeDados)

		recebido, _ := armazenamento.ObterPontuacaoJogador("Rafael")

		esperado := 10

		helpers.VerificaPontuacao(t, recebido, esperado)
	})

	t.Run("leitor de liga", func (t *testing.T) {
		conteudoArquivo := `[
			{"Nome": "Rafael", "Pontos": 10},
			{"Nome": "Vanessa", "Pontos": 20},
			{"Nome": "Pedro", "Pontos": 30}
		]`

		bancoDeDados, removerArquivo := helpers.CriarArquivoTemporario(t, conteudoArquivo)
		defer removerArquivo()

		armazenamento, _ := NovoArmazenamentoJogadorDoArquivo(bancoDeDados)

		recebido := armazenamento.ObterLiga()

		esperado := liga.Liga{
			{Nome: "Pedro", Pontos: 30},
			{Nome: "Vanessa", Pontos: 20},
			{Nome: "Rafael", Pontos: 10},
		}

		helpers.VerificaLiga(t, recebido, esperado)

		recebido = armazenamento.ObterLiga()

		helpers.VerificaLiga(t, recebido, esperado)
	})

	t.Run("registra vitoria e persiste no arquivo", func(t *testing.T) {
		conteudoArquivo := `[
			{"Nome": "Rafael", "Pontos": 10},
			{"Nome": "Vanessa", "Pontos": 20},
			{"Nome": "Pedro", "Pontos": 30}
		]`

		bancoDeDados, removerArquivo := helpers.CriarArquivoTemporario(t, conteudoArquivo)
		defer removerArquivo()

		armazenamento, _ := NovoArmazenamentoJogadorDoArquivo(bancoDeDados)

		armazenamento.RegistrarVitoria("Rafael")

		recebido, _ := armazenamento.ObterPontuacaoJogador("Rafael")

		esperado := 11

		helpers.VerificaPontuacao(t, recebido, esperado)
	})
}
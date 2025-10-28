package tape_test

import (
	"io"
	armazenamento "tdd/application/1-http/linha-de-comando/armazenamento"
	helpers "tdd/application/1-http/linha-de-comando/helpers"
	tape "tdd/application/1-http/linha-de-comando/tape"
	"testing"
)

func TestFitaEscrever(t *testing.T) {
	t.Run("escreve no arquivo a partir do inicio", func(t *testing.T) {
		arquivo, limpa := helpers.CriarArquivoTemporario(t, "ola mundo")
		defer limpa()

		fita := tape.NewFita(arquivo)

		fita.Write([]byte("ola"))

		arquivo.Seek(0, 0)
		conteudo, _ := io.ReadAll(arquivo)

		esperado := "ola mundo"
		obtido := string(conteudo)

		if obtido != esperado {
			t.Errorf("obtido '%s', esperado '%s'", obtido, esperado)
		}
	})

	t.Run("funciona com um arquivo vazio", func(t *testing.T) {
		arquivo, limpa := helpers.CriarArquivoTemporario(t, "")
		defer limpa()

		_, err := armazenamento.NovoArmazenamentoJogadorDoArquivo(arquivo)
		helpers.DefineSemErro(t, err)
	})
}

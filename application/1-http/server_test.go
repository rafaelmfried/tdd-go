package server_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	server "tdd/application/1-http"
	"testing"
)

type EsbocoArmazenamentoJogador struct {
	pontuacoes map[string]int
}

func (e *EsbocoArmazenamentoJogador) ObterPontuacaoJogador(nome string) int {
	pontuacao := e.pontuacoes[nome]
	return pontuacao
}

func TestObterJogadores(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{
		map[string]int{
			"Rafael": 20,
			"Vanessa": 15,
			"Pedro": 10,
		},
	}

	server := server.NewServidorJogador(&armazenamento)
	t.Run("retorna o resultdo de Rafael", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Rafael")
		resposta := httptest.NewRecorder()

		server.ServeHTTP(resposta, requisicao)

		obtido := resposta.Body.String()
		esperado := "20"
		status := resposta.Code
		statusEsperado := http.StatusOK

		verificarCorpoRequisicao(t, obtido, esperado)
		verificarStatusCodeRequisicao(t, status, statusEsperado)
	})

	t.Run("retorna o resultdo de Vanessa", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Vanessa")
		resposta := httptest.NewRecorder()

		server.ServeHTTP(resposta, requisicao)

		obtido := resposta.Body.String()
		esperado := "15"
		status := resposta.Code
		statusEsperado := http.StatusOK

		verificarCorpoRequisicao(t, obtido, esperado)
		verificarStatusCodeRequisicao(t, status, statusEsperado)
	})

	t.Run("jogador n encontrado erro 404 not found", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Marcos")
		resposta := httptest.NewRecorder()

		server.ServeHTTP(resposta, requisicao)

		status := resposta.Code
		statusEsperado := http.StatusNotFound

		if status != statusEsperado {
			t.Errorf("o resultdo obtido foi %d quando deveria ter sido: %d", status, statusEsperado)
		}
	})
}

func novaRequisicaoObterPontuacao(nome string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/jogadores/%s", nome), nil)
	return requisicao
}

func verificarCorpoRequisicao(t *testing.T, recebido, esperado string) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("obtido '%s', esperado '%s'", recebido, esperado)
	}
}

func verificarStatusCodeRequisicao(t *testing.T, recebido, esperado int) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("nao recebeu o codigo esperado: %d, recebido: %d", esperado, recebido)
	}
}
package server_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	server "tdd/application/1-http"
	"testing"
)

func TestObterJogadores(t *testing.T) {
	server := &server.ServidorJogador{}
	t.Run("retorna o resultdo de Rafael", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Rafael")
		resposta := httptest.NewRecorder()

		server.ServeHTTP(resposta, requisicao)

		obtido := resposta.Body.String()
		esperado := "20"

		verificarCorpoRequisicao(t, obtido, esperado)
	})

	t.Run("retorna o resultdo de Vanessa", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Vanessa")
		resposta := httptest.NewRecorder()

		server.ServeHTTP(resposta, requisicao)

		obtido := resposta.Body.String()
		esperado := "10"

		verificarCorpoRequisicao(t, obtido, esperado)
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
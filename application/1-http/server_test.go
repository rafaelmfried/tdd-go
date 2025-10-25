package server_test

import (
	"net/http"
	"net/http/httptest"
	server "tdd/application/1-http"
	"testing"
)

func TestObterJogadores(t *testing.T) {
	t.Run("retorna o resultdo de Rafael", func(t *testing.T) {
		requisicao, _ := http.NewRequest(http.MethodGet, "/jogadores/Rafael", nil)
		resposta := httptest.NewRecorder()

		server.ServidorJogador(resposta, requisicao)

		obtido := resposta.Body.String()
		esperado := "20"

		if obtido != esperado {
			t.Errorf("obtido '%s', esperado '%s'", obtido, esperado)
		}
	})

	t.Run("retorna o resultdo de Vanessa", func(t *testing.T) {
		requisicao, _ := http.NewRequest(http.MethodGet, "/jogadores/Vanessa", nil)
		resposta := httptest.NewRecorder()

		server.ServidorJogador(resposta, requisicao)

		obtido := resposta.Body.String()
		esperado := "10"

		if obtido != esperado {
			t.Errorf("obtido '%s', esperado '%s'", obtido, esperado)
		}
	})
}
/*
	Aqui a gente usou a biblioteca net/http/httptest para criar servidores HTTP de teste.
	Esses servidores simulam respostas lentas e rápidas para testar a função Corredor.
*/

package myselect_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	myselect "tdd/12-myselect"
	"testing"
	"time"
)

func TestCorredor(t *testing.T) {
	t.Run("teste de concorrencia com select", func(t *testing.T) {
		servidorLento := criarServidorComDelay(20 * time.Millisecond)
		servidorRapido := criarServidorComDelay(10 * time.Millisecond)

		// Faz chamar no final, porem mais facil de garantir o fechamento
		defer servidorLento.Close()
		defer servidorRapido.Close()

		URLLenta := servidorLento.URL
		URLRapida := servidorRapido.URL

		esperado := URLRapida

		obtido, err := myselect.Corredor(URLLenta, URLRapida)

		if err != nil {
			t.Errorf("esperado sucesso, mas obteve erro: %v", err)
		}

		if obtido != esperado {
			t.Errorf("obtido: %q, esperado: %q", obtido, esperado)
		}

		fmt.Printf("esperado: %q, obtido: %q", esperado, obtido)
	})

	t.Run("retorna um erro se o servidor demorar muito tempo", func(t *testing.T) {
		servidorLento := criarServidorComDelay(3 * time.Second)
		servidorRapido := criarServidorComDelay(2 * time.Second)

		// Faz chamar no final, porem mais facil de garantir o fechamento
		defer servidorLento.Close()
		defer servidorRapido.Close()

		URLLenta := servidorLento.URL
		URLRapida := servidorRapido.URL

		_, err := myselect.Corredor(URLLenta, URLRapida)

		if err == nil {
			t.Errorf("esperado um erro, mas obteve sucesso")
		}
	})
}

func criarServidorComDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
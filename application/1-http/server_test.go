package server_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	. "tdd/application/1-http"
	"testing"
)

type EsbocoArmazenamentoJogador struct {
	pontuacoes map[string]int
	registrosVitorias []string
}

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

func TestObterJogadores(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{
		pontuacoes: map[string]int{
			"Rafael": 20,
			"Vanessa": 15,
			"Pedro": 10,
		},
		registrosVitorias: []string{},
	}

	server := NewServidorJogador(&armazenamento)
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

func TestArmazenamentoVitorias(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{
		map[string]int{},
		nil,
	}

	server := NewServidorJogador(&armazenamento)

	t.Run("retorna status aceito para chamadas ao metodo POST", func(t *testing.T) {
		requisicao := novaRequisicaoRegistroVitoriaPost("Rafael")
		resposta := httptest.NewRecorder()

		server.ServeHTTP(resposta, requisicao)

		verificarStatusCodeRequisicao(t, resposta.Code, http.StatusAccepted)
		chamadasEsperadas := 1
		chamadasObtidas := len(armazenamento.registrosVitorias)
		if (chamadasObtidas != chamadasEsperadas) {
			t.Errorf("eram esperadas %d chamadas e obtivemos %d", chamadasEsperadas, chamadasObtidas)
		}
	})

	t.Run("registra vitorias na chamada ao metodo http post", func(t *testing.T) {
		jogador := "Vanessa"

		requisicao := novaRequisicaoRegistroVitoriaPost(jogador)
		resposta := httptest.NewRecorder()

		server.ServeHTTP(resposta, requisicao)

		chamadasEsperadas := 2
		chamadasObtidas := len(armazenamento.registrosVitorias)

		if chamadasObtidas != chamadasEsperadas {
			t.Errorf("eram esperadas %d chamadas e obtivemos %d", chamadasEsperadas, chamadasObtidas)
		}

		jogadorRecebido := armazenamento.registrosVitorias[chamadasEsperadas - 1]

		if jogadorRecebido != jogador {
			t.Errorf("nao registrou o vencedor corretamente, recebemos %s, esperava %s", jogadorRecebido, jogador)
		}
	})

	t.Run("registra e busca as vitorias", func(t *testing.T) {
		armazenamento := NovoArmazenamentoJogadorInMemory()
		servidor := NewServidorJogador(armazenamento)
		// investigar porque a pontuacao inicial esta como 0 para todos jogadores pelo novo armazenamento in memory
		jogador := "Marcos"

		requestPontuacaoInicial := httptest.NewRecorder()
		servidor.ServeHTTP(requestPontuacaoInicial, novaRequisicaoObterPontuacao(jogador))
		// converter string para int em go
		novo := requestPontuacaoInicial.Body.String()
		t.Logf("recebido no requesto pontuacao inicial %s", novo)
		pontuacaoInicial, _ := strconv.Atoi(novo)

		servidor.ServeHTTP(httptest.NewRecorder(), novaRequisicaoRegistroVitoriaPost(jogador))
		servidor.ServeHTTP(httptest.NewRecorder(), novaRequisicaoRegistroVitoriaPost(jogador))
		servidor.ServeHTTP(httptest.NewRecorder(), novaRequisicaoRegistroVitoriaPost(jogador))

		pontuacaoEsperada := pontuacaoInicial + 3

		t.Logf("pontuacao esperada: %d", pontuacaoEsperada)

		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, novaRequisicaoObterPontuacao(jogador))
		verificarStatusCodeRequisicao(t, resposta.Code, http.StatusOK)

		verificarCorpoRequisicao(t, resposta.Body.String(), strconv.Itoa(pontuacaoEsperada))

	})
}

func novaRequisicaoRegistroVitoriaPost(nome string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/jogadores/%s", nome), nil)
	return requisicao
}

func TestLiga(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{}
	servidor := NewServidorJogador(&armazenamento)

	t.Run("retorna 200 em liga", func(t *testing.T) {
		requisicao := novaRequisicaoBuscaLiga()
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificarStatusCodeRequisicao(t, resposta.Code, http.StatusOK)
	})
}

func novaRequisicaoBuscaLiga() *http.Request {
	requisicao, _ := http.NewRequest(http.MethodGet, "/liga", nil)
	return requisicao
}
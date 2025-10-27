// Refatorar para testes orientados a tabela (table-driven tests) e helpers de verificação

package server_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strconv"
	. "tdd/application/1-http"
	liga "tdd/application/1-http/liga"
	"testing"
)

type EsbocoArmazenamentoJogador struct {
	pontuacoes map[string]int
	registrosVitorias []string
	liga []liga.Jogador
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

func (e *EsbocoArmazenamentoJogador) ObterLiga() liga.Liga  {
	return e.liga
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

func TestArmazenamentoVitorias(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{
		map[string]int{},
		nil,
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
		bancoDeDados, limpaBancoDeDados := criarArquivoTemporario(t, "")
		defer limpaBancoDeDados()
		armazenamento := NovoArmazenamentoJogadorDoArquivo(bancoDeDados)
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

	t.Run("registra vitoria de novos jogadores", func(t *testing.T) {
		bancoDeDados, limpaBancoDeDados := criarArquivoTemporario(t, "")
		defer limpaBancoDeDados()

		armazenamento := NovoArmazenamentoJogadorDoArquivo(bancoDeDados)

		armazenamento.RegistrarVitoria("Pepperoni")

		recebido, _ := armazenamento.ObterPontuacaoJogador("Pepperoni")

		esperado := 1

		verificaPontuacao(t, recebido, esperado)
	})
}

func TestLiga(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{}
	servidor := NewServidorJogador(&armazenamento)

	t.Run("retorna 200 em liga", func(t *testing.T) {
		requisicao := novaRequisicaoBuscaLiga()
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		var obtido []liga.Jogador

		err := json.NewDecoder(resposta.Body).Decode(&obtido)

		if err != nil {
			t.Fatalf("nao foi possivel decodificar a resposta do servidor %v", err)
		}

		verificarStatusCodeRequisicao(t, resposta.Code, http.StatusOK)
	})

	t.Run("retorna a liga esperada como json", func(t *testing.T) {
		ligaEsperada := []liga.Jogador{
			{Nome: "Rafael", Pontos: 30},
			{Nome: "Vanessa", Pontos: 40},
			{Nome: "Pedro", Pontos: 20},
			{Nome: "Marcos", Pontos: 3},
		}

		armazenamento := EsbocoArmazenamentoJogador{
			liga: ligaEsperada,
		}

		servidor := NewServidorJogador(&armazenamento)

		requisicao := novaRequisicaoBuscaLiga()
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		obtido := obterLigaDaResposta(t, resposta.Body)

		verificarStatusCodeRequisicao(t, resposta.Code, http.StatusOK)

		verificaLiga(t, obtido, ligaEsperada)

		verificaContentType(t, resposta, JSONContentType)
	})
}

func TestSistemaDeArquivoDeArmazenamentoDoJogador(t *testing.T) {
	t.Run("liga eh carregada de um leitor", func(t *testing.T) {
		conteudoArquivo := `[
			{"Nome": "Rafael", "Pontos": 10},
			{"Nome": "Vanessa", "Pontos": 20},
			{"Nome": "Pedro", "Pontos": 30}
		]`

		bancoDeDados, limpaBancoDeDados := criarArquivoTemporario(t, conteudoArquivo)
		defer limpaBancoDeDados()

		armazenamento := NovoArmazenamentoJogadorDoArquivo(bancoDeDados)

		recebido := armazenamento.ObterLiga()

		esperado := []liga.Jogador{
			{Nome: "Rafael", Pontos: 10},
			{Nome: "Vanessa", Pontos: 20},
			{Nome: "Pedro", Pontos: 30},
		}

		verificaLiga(t, recebido, esperado)
	})

	t.Run("pegar pontuacao do jogador do arquivo", func(t *testing.T) {
		conteudoArquivo := `[
			{"Nome": "Rafael", "Pontos": 10},
			{"Nome": "Vanessa", "Pontos": 20},
			{"Nome": "Pedro", "Pontos": 30}
		]`
		bancoDeDados, limpaBancoDeDados := criarArquivoTemporario(t, conteudoArquivo)
		defer limpaBancoDeDados()

		armazenamento := NovoArmazenamentoJogadorDoArquivo(bancoDeDados)

		recebido, _ := armazenamento.ObterPontuacaoJogador("Rafael")

		esperado := 10

		verificaPontuacao(t, recebido, esperado)
	})

	t.Run("leitor de liga", func (t *testing.T) {
		conteudoArquivo := `[
			{"Nome": "Rafael", "Pontos": 10},
			{"Nome": "Vanessa", "Pontos": 20},
			{"Nome": "Pedro", "Pontos": 30}
		]`

		bancoDeDados, removerArquivo := criarArquivoTemporario(t, conteudoArquivo)
		defer removerArquivo()

		armazenamento := NovoArmazenamentoJogadorDoArquivo(bancoDeDados)

		recebido := armazenamento.ObterLiga()

		esperado := []liga.Jogador{
			{Nome: "Rafael", Pontos: 10},
			{Nome: "Vanessa", Pontos: 20},
			{Nome: "Pedro", Pontos: 30},
		}

		verificaLiga(t, recebido, esperado)

		recebido = armazenamento.ObterLiga()

		verificaLiga(t, recebido, esperado)
	})

	t.Run("registra vitoria e persiste no arquivo", func(t *testing.T) {
		conteudoArquivo := `[
			{"Nome": "Rafael", "Pontos": 10},
			{"Nome": "Vanessa", "Pontos": 20},
			{"Nome": "Pedro", "Pontos": 30}
		]`

		bancoDeDados, removerArquivo := criarArquivoTemporario(t, conteudoArquivo)
		defer removerArquivo()

		armazenamento := NovoArmazenamentoJogadorDoArquivo(bancoDeDados)

		armazenamento.RegistrarVitoria("Rafael")

		recebido, _ := armazenamento.ObterPontuacaoJogador("Rafael")

		esperado := 11

		verificaPontuacao(t, recebido, esperado)
	})
}

// Helpers

// Request helpers
func novaRequisicaoObterPontuacao(nome string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/jogadores/%s", nome), nil)
	return requisicao
}

func novaRequisicaoRegistroVitoriaPost(nome string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/jogadores/%s", nome), nil)
	return requisicao
}

func novaRequisicaoBuscaLiga() *http.Request {
	requisicao, _ := http.NewRequest(http.MethodGet, "/liga", nil)
	return requisicao
}

// Verification helpers
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

func verificaContentType(t *testing.T, resposta *httptest.ResponseRecorder, esperado string) {
	t.Helper()
	obtido := resposta.Header().Get("content-type")
	if obtido != esperado {
		t.Errorf("o content type obtido foi %s, esperado %s", obtido, esperado)
	}
}

func verificaLiga(t *testing.T, obtido, esperado []liga.Jogador) {
	t.Helper()
	fmt.Printf("VERIFICANDO LIGA: %v, %v", obtido, esperado)
	if !reflect.DeepEqual(obtido, esperado) {
		t.Errorf("obtido %v, esperado %v", obtido, esperado)
	}
}

func verificaPontuacao(t *testing.T, obtido, esperado int) {
	t.Helper()
	if obtido != esperado {
		t.Errorf("obtido %d, esperado %d", obtido, esperado)
	}
}


// Others
func obterLigaDaResposta(t *testing.T, body io.Reader) (liga []liga.Jogador) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&liga)

	if err != nil {
		panic(fmt.Sprintf("nao foi possivel decodificar a resposta do servidor %v", err))
	}

	return
}

func criarArquivoTemporario(t *testing.T, conteudo string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	arquivotmp, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("nao foi possivel criar um arquivo temporario %v", err)
	}

	arquivotmp.Write([]byte(conteudo))

	removeArquivo := func() {
		arquivotmp.Close()
		os.Remove(arquivotmp.Name())
	}

	return arquivotmp, removeArquivo
}
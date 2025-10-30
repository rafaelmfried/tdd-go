// Refatorar para testes orientados a tabela (table-driven tests) e helpers de verificação
package servidor_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	storage "tdd/application/1-http/linha-de-comando/armazenamento"
	"tdd/application/1-http/linha-de-comando/helpers"
	liga "tdd/application/1-http/linha-de-comando/liga"
	. "tdd/application/1-http/linha-de-comando/servidor"
	"testing"
)

func TestArmazenamentoVitorias(t *testing.T) {
	armazenamento := helpers.EsbocoArmazenamentoJogador{
		Pontuacoes:      map[string]int{},
		RegistrosVitorias: nil,
		Liga:           liga.Liga{},
	}

	server := NewServidorJogador(&armazenamento)

	t.Run("retorna status aceito para chamadas ao metodo POST", func(t *testing.T) {
		requisicao := helpers.NovaRequisicaoRegistroVitoriaPost("Rafael")
		resposta := httptest.NewRecorder()

		server.ServeHTTP(resposta, requisicao)

		helpers.VerificarStatusCodeRequisicao(t, resposta.Code, http.StatusAccepted)
		chamadasEsperadas := 1
		chamadasObtidas := len(armazenamento.RegistrosVitorias)
		if (chamadasObtidas != chamadasEsperadas) {
			t.Errorf("eram esperadas %d chamadas e obtivemos %d", chamadasEsperadas, chamadasObtidas)
		}
	})

	t.Run("registra vitorias na chamada ao metodo http post", func(t *testing.T) {
		jogador := "Vanessa"

		requisicao := helpers.NovaRequisicaoRegistroVitoriaPost(jogador)
		resposta := httptest.NewRecorder()

		server.ServeHTTP(resposta, requisicao)

		chamadasEsperadas := 2
		chamadasObtidas := len(armazenamento.RegistrosVitorias)

		if chamadasObtidas != chamadasEsperadas {
			t.Errorf("eram esperadas %d chamadas e obtivemos %d", chamadasEsperadas, chamadasObtidas)
		}

		jogadorRecebido := armazenamento.RegistrosVitorias[chamadasEsperadas - 1]

		if jogadorRecebido != jogador {
			t.Errorf("nao registrou o vencedor corretamente, recebemos %s, esperava %s", jogadorRecebido, jogador)
		}
	})

	t.Run("registra e busca as vitorias", func(t *testing.T) {
		bancoDeDados, limpaBancoDeDados := helpers.CriarArquivoTemporario(t, "[]")
		defer limpaBancoDeDados()
		armazenamento, err := storage.NovoArmazenamentoJogadorDoArquivo(bancoDeDados)
		if err != nil {
			t.Fatalf("nao foi possivel criar o armazenamento do jogador a partir do arquivo: %v", err)
		}
		servidor := NewServidorJogador(armazenamento)
		// investigar porque a pontuacao inicial esta como 0 para todos jogadores pelo novo armazenamento in memory
		jogador := "Marcos"

		requestPontuacaoInicial := httptest.NewRecorder()
		servidor.ServeHTTP(requestPontuacaoInicial, helpers.NovaRequisicaoObterPontuacao(jogador))
		// converter string para int em go
		novo := requestPontuacaoInicial.Body.String()
		t.Logf("recebido no requesto pontuacao inicial %s", novo)
		pontuacaoInicial, _ := strconv.Atoi(novo)

		servidor.ServeHTTP(httptest.NewRecorder(), helpers.NovaRequisicaoRegistroVitoriaPost(jogador))
		servidor.ServeHTTP(httptest.NewRecorder(), helpers.NovaRequisicaoRegistroVitoriaPost(jogador))
		servidor.ServeHTTP(httptest.NewRecorder(), helpers.NovaRequisicaoRegistroVitoriaPost(jogador))

		pontuacaoEsperada := pontuacaoInicial + 3

		t.Logf("pontuacao esperada: %d", pontuacaoEsperada)

		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, helpers.NovaRequisicaoObterPontuacao(jogador))
		helpers.VerificarStatusCodeRequisicao(t, resposta.Code, http.StatusOK)

		helpers.VerificarCorpoRequisicao(t, resposta.Body.String(), strconv.Itoa(pontuacaoEsperada))
	})

	t.Run("registra vitoria de novos jogadores", func(t *testing.T) {
		bancoDeDados, limpaBancoDeDados := helpers.CriarArquivoTemporario(t, "[]")
		defer limpaBancoDeDados()

		armazenamento, _ := storage.NovoArmazenamentoJogadorDoArquivo(bancoDeDados)

		armazenamento.RegistrarVitoria("Pepperoni")

		recebido, _ := armazenamento.ObterPontuacaoJogador("Pepperoni")

		esperado := 1

		helpers.VerificaPontuacao(t, recebido, esperado)
	})
}


func TestObterJogadores(t *testing.T) {
	storage := helpers.EsbocoArmazenamentoJogador{
		Pontuacoes: map[string]int{
			"Rafael": 20,
			"Vanessa": 15,
			"Pedro": 10,
		},
		RegistrosVitorias: []string{},
	}

	server := NewServidorJogador(&storage)
	t.Run("retorna o resultdo de Rafael", func(t *testing.T) {
		requisicao := helpers.NovaRequisicaoObterPontuacao("Rafael")
		resposta := httptest.NewRecorder()

		server.ServeHTTP(resposta, requisicao)

		obtido := resposta.Body.String()
		esperado := "20"
		status := resposta.Code
		statusEsperado := http.StatusOK

		helpers.VerificarCorpoRequisicao(t, obtido, esperado)
		helpers.VerificarStatusCodeRequisicao(t, status, statusEsperado)
	})

	t.Run("retorna o resultdo de Vanessa", func(t *testing.T) {
		requisicao := helpers.NovaRequisicaoObterPontuacao("Vanessa")
		resposta := httptest.NewRecorder()

		server.ServeHTTP(resposta, requisicao)

		obtido := resposta.Body.String()
		esperado := "15"
		status := resposta.Code
		statusEsperado := http.StatusOK

		helpers.VerificarCorpoRequisicao(t, obtido, esperado)
		helpers.VerificarStatusCodeRequisicao(t, status, statusEsperado)
	})

	t.Run("jogador n encontrado erro 404 not found", func(t *testing.T) {
		requisicao := helpers.NovaRequisicaoObterPontuacao("Marcos")
		resposta := httptest.NewRecorder()

		server.ServeHTTP(resposta, requisicao)

		status := resposta.Code
		statusEsperado := http.StatusNotFound

		if status != statusEsperado {
			t.Errorf("o resultdo obtido foi %d quando deveria ter sido: %d", status, statusEsperado)
		}
	})
}

func TestLiga(t *testing.T) {
	s := helpers.EsbocoArmazenamentoJogador{}
	servidor := NewServidorJogador(&s)

	t.Run("retorna 200 em liga", func(t *testing.T) {
		requisicao := helpers.NovaRequisicaoBuscaLiga()
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		var obtido []liga.Jogador

		err := json.NewDecoder(resposta.Body).Decode(&obtido)

		if err != nil {
			t.Fatalf("nao foi possivel decodificar a resposta do servidor %v", err)
		}

		helpers.VerificarStatusCodeRequisicao(t, resposta.Code, http.StatusOK)
	})

	t.Run("retorna a liga esperada como json", func(t *testing.T) {
		ligaEsperada := []liga.Jogador{
			{Nome: "Rafael", Pontos: 30},
			{Nome: "Vanessa", Pontos: 40},
			{Nome: "Pedro", Pontos: 20},
			{Nome: "Marcos", Pontos: 3},
		}

		armazenamento := helpers.EsbocoArmazenamentoJogador{
			Liga: ligaEsperada,
		}

		servidor := NewServidorJogador(&armazenamento)

		requisicao := helpers.NovaRequisicaoBuscaLiga()
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		obtido := helpers.ObterLigaDaResposta(t, resposta.Body)

		helpers.VerificarStatusCodeRequisicao(t, resposta.Code, http.StatusOK)

		helpers.VerificaLiga(t, obtido, ligaEsperada)

		helpers.VerificaContentType(t, resposta, JSONContentType)
	})

	t.Run("liga eh retornada em ordem decrescente", func(t *testing.T) {
		ligaString := `[
			{"Nome": "Rafael", "Pontos": 30},
			{"Nome": "Vanessa", "Pontos": 40},
			{"Nome": "Pedro", "Pontos": 20},
			{"Nome": "Marcos", "Pontos": 3}]`

		bancoDeDados, limpaBancoDeDados := helpers.CriarArquivoTemporario(t, ligaString)

		defer limpaBancoDeDados()

		armazenamento, _ := storage.NovoArmazenamentoJogadorDoArquivo(bancoDeDados)

		recebido := armazenamento.ObterLiga()

		ligaEsperada := liga.Liga{
			{Nome: "Vanessa", Pontos: 40},
			{Nome: "Rafael", Pontos: 30},
			{Nome: "Pedro", Pontos: 20},
			{Nome: "Marcos", Pontos: 3},
		}

		helpers.VerificaLiga(t, recebido, ligaEsperada)
		recebido = armazenamento.ObterLiga()

		helpers.VerificaLiga(t, recebido, ligaEsperada)
	})

	t.Run("GET /jogo retorna 200 OK", func(t *testing.T) {
		servidor := NewServidorJogador(&helpers.EsbocoArmazenamentoJogador{})

		requisicao := helpers.NovaRequisicaoJogo()
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		helpers.VerificarStatusCodeRequisicao(t, resposta.Code, http.StatusOK)
	})
}
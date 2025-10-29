package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"tdd/application/1-http/linha-de-comando/armazenamento"
	"tdd/application/1-http/linha-de-comando/liga"
	"testing"
)

type EsbocoArmazenamentoJogador struct {
	Pontuacoes map[string]int
	RegistrosVitorias []string
	Liga liga.Liga
}

var ErrJogadorNotFound = armazenamento.ErrJogadorNotFound

func (e *EsbocoArmazenamentoJogador) ObterPontuacaoJogador(nome string) (pontuacao int, err error) {
	pontuacao = e.Pontuacoes[nome]
	if pontuacao == 0 {
		return 0, ErrJogadorNotFound
	}
	return pontuacao, nil
}

func (e *EsbocoArmazenamentoJogador) RegistrarVitoria(nome string) {
	e.RegistrosVitorias = append(e.RegistrosVitorias, nome)
}

func (e *EsbocoArmazenamentoJogador) ObterLiga() liga.Liga  {
	return e.Liga
}

// Request helpers
func NovaRequisicaoObterPontuacao(nome string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/jogadores/%s", nome), nil)
	return requisicao
}

func NovaRequisicaoRegistroVitoriaPost(nome string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/jogadores/%s", nome), nil)
	return requisicao
}

func NovaRequisicaoBuscaLiga() *http.Request {
	requisicao, _ := http.NewRequest(http.MethodGet, "/liga", nil)
	return requisicao
}

// Verification helpers
func VerificarCorpoRequisicao(t *testing.T, recebido, esperado string) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("obtido '%s', esperado '%s'", recebido, esperado)
	}
}

func VerificarStatusCodeRequisicao(t *testing.T, recebido, esperado int) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("nao recebeu o codigo esperado: %d, recebido: %d", esperado, recebido)
	}
}

func VerificaContentType(t *testing.T, resposta *httptest.ResponseRecorder, esperado string) {
	t.Helper()
	obtido := resposta.Header().Get("content-type")
	if obtido != esperado {
		t.Errorf("o content type obtido foi %s, esperado %s", obtido, esperado)
	}
}

func VerificaLiga(t *testing.T, obtido, esperado liga.Liga) {
	t.Helper()
	fmt.Printf("VERIFICANDO LIGA: %v, %v", obtido, esperado)
	if !reflect.DeepEqual(obtido, esperado) {
		t.Errorf("obtido %v, esperado %v", obtido, esperado)
	}
}

func VerificaPontuacao(t *testing.T, obtido, esperado int) {
	t.Helper()
	if obtido != esperado {
		t.Errorf("obtido %d, esperado %d", obtido, esperado)
	}
}


// Others
func ObterLigaDaResposta(t *testing.T, body io.Reader) (liga liga.Liga) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&liga)

	if err != nil {
		panic(fmt.Sprintf("nao foi possivel decodificar a resposta do servidor %v", err))
	}

	return
}

func CriarArquivoTemporario(t *testing.T, conteudo string) (*os.File, func()) {
	t.Helper()

	arquivotmp, err := os.CreateTemp("", "db")

	
	if err != nil {
		t.Fatalf("nao foi possivel criar um arquivo temporario %v", err)
	}
	
	arquivotmp.Write([]byte(conteudo))

	arquivotmp.Seek(0, 0)

	removeArquivo := func() {
		arquivotmp.Close()
		os.Remove(arquivotmp.Name())
	}

	return arquivotmp, removeArquivo
}

func DefineSemErro(t *testing.T, err error) {
    t.Helper()
    if err != nil {
        t.Fatalf("não esperava um erro mas obteve um, %v", err)
    }
}

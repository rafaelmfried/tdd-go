package servidor

import (
	"encoding/json"
	"fmt"
	"net/http"
	armazenamento "tdd/application/1-http/linha-de-comando/armazenamento"
	"tdd/application/1-http/linha-de-comando/liga"
)

const JSONContentType = "application/json"

type Jogador = liga.Jogador

var ErrJogadorNotFound = armazenamento.ErrJogadorNotFound

type ServidorJogador struct {
	armazenamento armazenamento.ArmazenamentoJogador
	http.Handler
}

func NewServidorJogador(armazenamento armazenamento.ArmazenamentoJogador) *ServidorJogador {
	s := new(ServidorJogador)
	s.armazenamento = armazenamento
	roteador := http.NewServeMux()
	roteador.Handle("/jogadores/", http.HandlerFunc(s.tratarRequisicaoJogador))
	roteador.Handle("/liga", http.HandlerFunc(s.tratarRequisicaoLiga))
	s.Handler = roteador
	return s
}

func (s *ServidorJogador) registrarVitoria(writer http.ResponseWriter, request *http.Request) {
	jogador := request.URL.Path[len("/jogadores/"):]
	s.armazenamento.RegistrarVitoria(jogador)
	writer.WriteHeader(http.StatusAccepted)
} 

func (s *ServidorJogador) mostrarPontuacao(writer http.ResponseWriter, request http.Request) {
	jogador := request.URL.Path[len("/jogadores/"):]
	pontuacao, err := s.armazenamento.ObterPontuacaoJogador(jogador)
	if err == ErrJogadorNotFound {
		writer.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(writer, pontuacao)
}

func (s *ServidorJogador) manipulaLiga(writer http.ResponseWriter, _ http.Request) {
	tabelaLiga := s.armazenamento.ObterLiga()
	fmt.Printf("tabela liga: %v", tabelaLiga)
	writer.Header().Set("content-type", JSONContentType)
	json.NewEncoder(writer).Encode(tabelaLiga)

	writer.WriteHeader(http.StatusOK)
}

func (s *ServidorJogador) tratarRequisicaoJogador(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		s.registrarVitoria(writer, request)
	case http.MethodGet:
		s.mostrarPontuacao(writer, *request)
	}
}

func (s *ServidorJogador) tratarRequisicaoLiga(writer http.ResponseWriter, request *http.Request) {
	s.manipulaLiga(writer, *request)
}

func Server() {
}
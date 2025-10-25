package server

import (
	"fmt"
	"log"
	"net/http"
)

type Jogador struct {
	Nome  string
	Pontos int
}

var jogadores = map[string]Jogador{
	"Rafael":  {Nome: "Rafael", Pontos: 30},
	"Vanessa": {Nome: "Vanessa", Pontos: 40},
	"Pedro":   {Nome: "Pedro", Pontos: 20},
}

var vitorias []string

var ErrJogadorNotFound = fmt.Errorf("Jogador n encontrado")

type ArmazenamentoJogador interface {
	ObterPontuacaoJogador(nome string) (pontuacao int, err error)
	RegistrarVitoria(nome string)
}

type ArmazenamentoJogadorInMemory struct {
	storage map[string]int
}

func NovoArmazenamentoJogadorInMemory() *ArmazenamentoJogadorInMemory {
	return &ArmazenamentoJogadorInMemory{map[string]int{}}
}

func (a *ArmazenamentoJogadorInMemory) ObterPontuacaoJogador(nome string) (pontuacao int, err error) {
	return obterPontuacaoJogador(nome)
}

func (a *ArmazenamentoJogadorInMemory) RegistrarVitoria(nome string) {
	registraVitoria(nome)
}

type ServidorJogador struct {
	armazenamento ArmazenamentoJogador
	roteador      *http.ServeMux
}

func NewServidorJogador(armazenamento ArmazenamentoJogador) *ServidorJogador {
	roteador := http.NewServeMux()
	servidor := &ServidorJogador{armazenamento: armazenamento, roteador: roteador}

	roteador.HandleFunc("/jogadores/", servidor.tratarRequisicaoJogador)
	roteador.HandleFunc("/liga", servidor.tratarRequisicaoLiga)

	return servidor
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

func (s *ServidorJogador) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.roteador.ServeHTTP(writer, request)
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
	writer.WriteHeader(http.StatusOK)
}

func obterPontuacaoJogador(nome string) (pontuacao int, err error) {
	if jogador, ok := jogadores[nome]; ok {
		return jogador.Pontos, nil
	}
	return 0, ErrJogadorNotFound
}

func registraVitoria(nome string) {
	vitorias = append(vitorias, nome)
	if jogador, ok := jogadores[nome]; ok {
		jogadores[nome] = Jogador{Nome: nome, Pontos: jogador.Pontos + 1}
		return
	}
	jogadores[nome] = Jogador{ Nome: nome, Pontos: 1}
}

func Server() {
	armazenamento := &ArmazenamentoJogadorInMemory{}
	handler := NewServidorJogador(armazenamento)
	tratador := http.HandlerFunc(handler.ServeHTTP)
	if err := http.ListenAndServe(":5324", tratador); err != nil {
		log.Fatalf("nao foi possivel escutar a porta 5324 %v", err)
	}
}
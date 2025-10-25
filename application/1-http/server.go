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

var ErrJogadorNotFound = fmt.Errorf("Jogador n encontrado")

type ArmazenamentoJogador interface {
	ObterPontuacaoJogador(nome string) int
}

type ArmazenamentoJogadorInMemory struct {}

func (a *ArmazenamentoJogadorInMemory) ObterPontuacaoJogador(nome string) int {
	pontuacao, err := obterPontuacaoJogador(nome)
	if err != nil {
		return 0
	}
	return pontuacao
}

type ServidorJogador struct {
	armazenamento ArmazenamentoJogador
}

func NewServidorJogador(armazenamento ArmazenamentoJogador) *ServidorJogador {
	return &ServidorJogador{armazenamento: armazenamento}
}

func (s *ServidorJogador) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	jogador := request.URL.Path[len("/jogadores/"):]
	writer.WriteHeader(http.StatusNotFound)
	fmt.Fprint(writer, s.armazenamento.ObterPontuacaoJogador(jogador))
}

func obterPontuacaoJogador(nome string) (pontuacao int, err error) {
	if jogador, ok := jogadores[nome]; ok {
		return jogador.Pontos, nil
	}
	return 0, ErrJogadorNotFound
}

func Server() {
	armazenamento := &ArmazenamentoJogadorInMemory{}
	handler := NewServidorJogador(armazenamento)
	tratador := http.HandlerFunc(handler.ServeHTTP)
	if err := http.ListenAndServe(":5324", tratador); err != nil {
		log.Fatalf("nao foi possivel escutar a porta 5324 %v", err)
	}
}
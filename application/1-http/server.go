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
	"Rafael":  {Nome: "Rafael", Pontos: 20},
	"Vanessa": {Nome: "Vanessa", Pontos: 10},
	"Pedro":   {Nome: "Pedro", Pontos: 15},
}

var ErrJogadorNotFound = fmt.Errorf("Jogador n encontrado")

type ArmazenamentoJogador interface {
	ObterPontuacaoJogador(nome string) int
}

type ServidorJogador struct {
	armazenamento ArmazenamentoJogador
}

func (s *ServidorJogador) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	jogador := request.URL.Path[len("/jogadores/"):]
	fmt.Fprint(writer, s.armazenamento.ObterPontuacaoJogador(jogador))
}

func ObterPontuacaoJogador(nome string) (pontuacao int, err error) {
	if jogador, ok := jogadores[nome]; ok {
		return jogador.Pontos, nil
	}
	return 0, ErrJogadorNotFound
}

func Server() {
	handler := &ServidorJogador{}
	tratador := http.HandlerFunc(handler.ServeHTTP)
	if err := http.ListenAndServe(":5324", tratador); err != nil {
		log.Fatalf("nao foi possivel escutar a porta 5324 %v", err)
	}
}
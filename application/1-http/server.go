package server

import (
	"fmt"
	"net/http"
)

type Jogador struct {
	Nome  string
	Pontos int
}

var jogadores = map[string]Jogador{
	"Rafael":  {Nome: "Rafael", Pontos: 20},
	"Vanessa": {Nome: "Vanessa", Pontos: 10},
}

func ServidorJogador(writer http.ResponseWriter, request *http.Request) {
	nomeJogador := request.URL.Path[len("/jogadores/"):]
	if jogador, ok := jogadores[nomeJogador]; ok {
		fmt.Fprint(writer, jogador.Pontos)
		return
	}
}
package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	liga "tdd/application/1-http/liga"
)

const JSONContentType = "application/json"

var jogadores = map[string]liga.Jogador{
	"Rafael":  {Nome: "Rafael", Pontos: 30},
	"Vanessa": {Nome: "Vanessa", Pontos: 40},
	"Pedro":   {Nome: "Pedro", Pontos: 20},
}

var vitorias []string

var ErrJogadorNotFound = fmt.Errorf("jogador n encontrado")

type ArmazenamentoJogador interface {
	ObterPontuacaoJogador(nome string) (pontuacao int, err error)
	RegistrarVitoria(nome string)
	ObterLiga() liga.Liga
}

type ArmazenamentoJogadorInMemory struct {
	storage map[string]int
}

type ArmazenamentoJogadorDoArquivo struct {
	bancoDeDados io.ReadWriteSeeker
}

func NovoArmazenamentoJogadorDoArquivo(bancoDeDados io.ReadWriteSeeker) *ArmazenamentoJogadorDoArquivo {
	return &ArmazenamentoJogadorDoArquivo{bancoDeDados: bancoDeDados}
}

func (f *ArmazenamentoJogadorDoArquivo) ObterLiga() liga.Liga {
	f.bancoDeDados.Seek(0, 0)
	liga, _ := liga.NovaLiga(f.bancoDeDados)
	return liga
}

func (f *ArmazenamentoJogadorDoArquivo) ObterPontuacaoJogador(nome string) int {
	jogador := f.ObterLiga().Find(nome)
	if jogador != nil {
		return jogador.Pontos
	}
	return 0
}

func (f *ArmazenamentoJogadorDoArquivo) SalvarVitoria(nome string) {
	liga := f.ObterLiga()

	for i, jogador := range liga {
		if jogador.Nome == nome {
			liga[i].Pontos++
		}
	}

	f.bancoDeDados.Seek(0, 0)
	json.NewEncoder(f.bancoDeDados).Encode(liga)
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

func (a *ArmazenamentoJogadorInMemory) ObterLiga() liga.Liga {
	return obterTabelaLiga()
}
type ServidorJogador struct {
	armazenamento ArmazenamentoJogador
	http.Handler
}

func NewServidorJogador(armazenamento ArmazenamentoJogador) *ServidorJogador {
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
	tabelaLiga := obterTabelaLiga()
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

func obterPontuacaoJogador(nome string) (pontuacao int, err error) {
	if jogador, ok := jogadores[nome]; ok {
		return jogador.Pontos, nil
	}
	return 0, ErrJogadorNotFound
}

func obterTabelaLiga() []liga.Jogador {
	reader := MapParaReader(jogadores)
	liga, _ := liga.NovaLiga(reader)

	return liga
}

func MapParaReader(jogadores map[string]liga.Jogador) io.Reader {
	var liga []liga.Jogador
	for _, jogador := range jogadores {
		liga = append(liga, jogador)
	}

	jsonData, err := json.Marshal(liga)
	if err != nil {
		log.Fatalf("nao foi possivel converter jogadores para JSON %v", err)
	}
	return bytes.NewReader(jsonData)
}

func registraVitoria(nome string) {
	vitorias = append(vitorias, nome)
	if jogador, ok := jogadores[nome]; ok {
		jogadores[nome] = liga.Jogador{Nome: nome, Pontos: jogador.Pontos + 1}
		return
	}
	jogadores[nome] = liga.Jogador{Nome: nome, Pontos: 1}
}

func Server() {
	// armazenamento := &ArmazenamentoJogadorDoArquivo{}
	// handler := NewServidorJogador(armazenamento)
	// tratador := http.HandlerFunc(handler.ServeHTTP)
	// if err := http.ListenAndServe(":5324", tratador); err != nil {
	// 	log.Fatalf("nao foi possivel escutar a porta 5324 %v", err)
	// }
}
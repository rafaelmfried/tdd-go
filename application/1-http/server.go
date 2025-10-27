package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"tdd/application/1-http/liga"
)

const JSONContentType = "application/json"

type Jogador = liga.Jogador

var ErrJogadorNotFound = fmt.Errorf("jogador n encontrado")

type ArmazenamentoJogador interface {
	ObterPontuacaoJogador(nome string) (pontuacao int, err error)
	RegistrarVitoria(nome string)
	ObterLiga() liga.Liga
}
type ArmazenamentoJogadorDoArquivo struct {
	bancoDeDados *json.Encoder
	liga liga.Liga
}

func iniciaArquivoDBJogador(arquivo *os.File) error {
	arquivo.Seek(0, 0)

	info, err := arquivo.Stat()

	if err != nil {
		return fmt.Errorf("nao foi possivel obter info do arquivo %s %v", arquivo.Name(), err)
	}

	if info.Size() == 0 {
		arquivo.Write([]byte("[]"))
		arquivo.Seek(0, 0)
	}
	return nil
}

func NovoArmazenamentoJogadorDoArquivo(arquivo *os.File) (*ArmazenamentoJogadorDoArquivo, error) {
	err := iniciaArquivoDBJogador(arquivo)
	
	if err != nil {
		return nil, fmt.Errorf("nao foi possivel iniciar o arquivo %s %v", arquivo.Name(), err)
	}

	liga, err := liga.NovaLiga(arquivo)

	if err != nil {
		return nil, fmt.Errorf("nao foi possivel criar a liga a partir do arquivo %v", err)
	}
	fita := NewFita(arquivo)
	return &ArmazenamentoJogadorDoArquivo{
		bancoDeDados: json.NewEncoder(fita),
		liga: liga,
	}, nil
}

func (f *ArmazenamentoJogadorDoArquivo) ObterLiga() liga.Liga {
	return f.liga
}

func (f *ArmazenamentoJogadorDoArquivo) ObterPontuacaoJogador(nome string) (int, error) {
	jogador := f.liga.Find(nome)

	fmt.Printf("JOGADOR: %v", jogador)
	if jogador != nil {
		fmt.Printf("PONTOS: %d", jogador.Pontos)
		return jogador.Pontos, nil
	}
	return 0, ErrJogadorNotFound
}

func (f *ArmazenamentoJogadorDoArquivo) RegistrarVitoria(nome string) {
	jogador := f.liga.Find(nome)
	fmt.Printf("SALVANDO JOGADOR: %v", jogador)
	if jogador != nil {
		fmt.Printf("SALVANDO JOGADOR EXISTE: %d", jogador.Pontos)
		jogador.Pontos++
	} else {
		f.liga = append(f.liga, liga.Jogador{Nome: nome, Pontos: 1})
		fmt.Printf("SALVANDO JOGADOR N EXISTE: %v", f.liga)
	}
	f.bancoDeDados.Encode(f.liga)
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

func MapParaReader(jogadores map[string]Jogador) io.ReadSeeker {
	var liga []Jogador
	for _, jogador := range jogadores {
		liga = append(liga, jogador)
	}

	jsonData, err := json.Marshal(liga)
	if err != nil {
		log.Fatalf("nao foi possivel converter jogadores para JSON %v", err)
	}
	return bytes.NewReader(jsonData)
}

func Server() {
	// armazenamento := &ArmazenamentoJogadorDoArquivo{}
	// handler := NewServidorJogador(armazenamento)
	// tratador := http.HandlerFunc(handler.ServeHTTP)
	// if err := http.ListenAndServe(":5324", tratador); err != nil {
	// 	log.Fatalf("nao foi possivel escutar a porta 5324 %v", err)
	// }
}
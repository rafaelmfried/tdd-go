package armazenamento

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"tdd/application/1-http/linha-de-comando/liga"
	"tdd/application/1-http/linha-de-comando/tape"
)

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
	fita := tape.NewFita(arquivo)
	return &ArmazenamentoJogadorDoArquivo{
		bancoDeDados: json.NewEncoder(fita),
		liga: liga,
	}, nil
}

func (f *ArmazenamentoJogadorDoArquivo) ObterLiga() liga.Liga {
	sort.Slice(f.liga, func(i, j int) bool {
		return f.liga[i].Pontos > f.liga[j].Pontos
	})
	return f.liga
}

func (f *ArmazenamentoJogadorDoArquivo) ObterPontuacaoJogador(nome string) (int, error) {
	jogador := f.liga.Find(nome)

	if jogador != nil {
		fmt.Printf("PONTOS: %d", jogador.Pontos)
		return jogador.Pontos, nil
	}
	return 0, ErrJogadorNotFound
}

func (f *ArmazenamentoJogadorDoArquivo) RegistrarVitoria(nome string) {
	jogador := f.liga.Find(nome)
	if jogador != nil {
		jogador.Pontos++
	} else {
		f.liga = append(f.liga, liga.Jogador{Nome: nome, Pontos: 1})
	}
	f.bancoDeDados.Encode(f.liga)
}
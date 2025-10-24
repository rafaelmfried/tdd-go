package maps

import (
	"fmt"
)

var ErroChaveInexistente = fmt.Errorf("chave inexistente")
var ErroChaveExistente = fmt.Errorf("chave já existente")
var ErroNaoEncontrado = fmt.Errorf("não encontrado")

type Dicionario struct {
	mapping map[string]string
}

func NewDicionario() *Dicionario {
	return &Dicionario{mapping: make(map[string]string)}
}

func (d Dicionario) Busca(chave string) (string, error) {
	if valor, existe := d.mapping[chave]; existe {
		return valor, nil
	}
	return "", ErroChaveInexistente
}

func (d Dicionario) Adicionar(chave, valor string) error {
	_, err := d.Busca(chave)

	if err == nil {
		return ErroChaveExistente
	}

	d.mapping[chave] = valor

	return nil
}

func (d Dicionario) Atualizar(chave, valor string) error {
	_, err := d.Busca(chave)

	if err != nil {
		return ErroNaoEncontrado
	}

	d.mapping[chave] = valor

	return nil
}

// funcao nativa do go para deletar um item de um map
func (d Dicionario) Deletar(chave string) {
	delete(d.mapping, chave)
}
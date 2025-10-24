package maps

import "fmt"

var ErroChaveInexistente = fmt.Errorf("chave inexistente")

type Dicionario map[string]string

func (d Dicionario) Busca(chave string) (string, error) {
	if valor, existe := d[chave]; existe {
		return valor, nil
	}
	return "", ErroChaveInexistente
}
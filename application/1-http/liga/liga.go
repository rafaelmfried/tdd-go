package liga

import (
	"encoding/json"
	"fmt"
	"io"
)

type Jogador struct {
	Nome   string
	Pontos int
}

func NovaLiga(reader io.Reader) ([]Jogador, error) {
	var liga []Jogador
	err := json.NewDecoder(reader).Decode(&liga)

	if err != nil {
		err = fmt.Errorf("problema ao parsear a liga, %v", err)
		return nil, err
	}
	return liga, nil
}
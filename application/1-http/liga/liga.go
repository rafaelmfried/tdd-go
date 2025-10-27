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

type Liga []Jogador

func NovaLiga(reader io.ReadSeeker) (Liga, error) {
	var liga Liga
	err := json.NewDecoder(reader).Decode(&liga)

	if err != nil {
		err = fmt.Errorf("problema ao parsear a liga, %v", err)
		return nil, err
	}
	return liga, nil
}

func (l Liga) Find(nome string) *Jogador {
	for i := range l {
		if l[i].Nome == nome {
			return &l[i]
		}
	}
	return nil
}
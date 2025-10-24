package maps_test

import (
	maps "tdd/8-maps"
	"testing"
)

func comparaErro(t *testing.T, resultado, esperado error) {
	t.Helper()
	if resultado == nil {
		t.Fatal("esperava um erro, mas não ocorreu nenhum")
	}
	if resultado != esperado {
		t.Errorf("resultado de erro '%s', esperado '%s'", resultado, esperado)
	}
}

func erroInexistente(t *testing.T, resultado error) {
	t.Helper()
	if resultado != nil {
		t.Fatal("esperava n existir erro, mas ocorreu um")
	}
}

func TestBusca(t *testing.T) {
	t.Run("Busca um valor existente em um mapa", func(t *testing.T) {
		dicionario := maps.Dicionario{"teste": "isso é um teste"}
		esperado := "isso é um teste"
		resultado, erro := dicionario.Busca("teste")

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}

		erroInexistente(t, erro)
	})
	
	t.Run("Retorna um erro ao buscar uma chave inexistente", func(t *testing.T) {
		dicionario := maps.Dicionario{}
		_, err := dicionario.Busca("inexistente")
		comparaErro(t, err, maps.ErroChaveInexistente)
	})
}
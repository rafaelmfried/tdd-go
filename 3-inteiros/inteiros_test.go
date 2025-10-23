package inteiros_test

import (
	inteiros "tdd/3-inteiros"
	"testing"
)

func TestAdicionar(t *testing.T) {
	t.Run("deve retornar a soma de dois inteiros positivos", func(t *testing.T) {
		resultado := inteiros.Adicionar(2, 3)
		valorEsperado := 5

		if resultado != valorEsperado {
			t.Errorf("resultado %d, valor esperado %d", resultado, valorEsperado)
		}
	})

	t.Run("deve retornar a soma de um inteiro positivo e um negativo", func(t *testing.T) {
		resultado := inteiros.Adicionar(5, -2)
		valorEsperado := 3

		if resultado != valorEsperado {
			t.Errorf("resultado %d, valor esperado %d", resultado, valorEsperado)
		}
	})

	t.Run("deve retornar a soma de dois inteiros negativos", func(t *testing.T) {
		resultado := inteiros.Adicionar(-4, -6)
		valorEsperado := -10

		if resultado != valorEsperado {
			t.Errorf("resultado %d, valor esperado %d", resultado, valorEsperado)
		}
	})
}
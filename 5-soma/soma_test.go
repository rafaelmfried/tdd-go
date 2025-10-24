package soma_test

import (
	"reflect"
	soma "tdd/5-soma"
	"testing"
)

func TestSoma(t *testing.T) {
	t.Run("soma de um slice com vários números", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}
		resultado := soma.Soma(numeros)
		esperado := 15

		if resultado != esperado {
			t.Errorf("resultado %d, esperado %d", resultado, esperado)
		}
	})

	t.Run("soma de um slice vazio", func(t *testing.T) {
		numeros := []int{}
		resultado := soma.Soma(numeros)
		esperado := 0

		if resultado != esperado {
			t.Errorf("resultado %d, esperado %d", resultado, esperado)
		}
	})

	t.Run("soma de um slice com números negativos", func(t *testing.T) {
		numeros := []int{-1, -2, -3, -4, -5}
		resultado := soma.Soma(numeros)
		esperado := -15

		if resultado != esperado {
			t.Errorf("resultado %d, esperado %d", resultado, esperado)
		}
	})

	t.Run("soma de um slice com números positivos e negativos", func(t *testing.T) {
		numeros := []int{-1, 2, -3, 4, -5}
		resultado := soma.Soma(numeros)
		esperado := -3

		if resultado != esperado {
			t.Errorf("resultado %d, esperado %d", resultado, esperado)
		}
	})
}

func TestSomaTudo(t *testing.T) {
	t.Run("soma de múltiplos slices", func(t *testing.T) {
		slice1 := []int{1, 2}
		slice2 := []int{3, 4}
		slice3 := []int{5, 6}
		resultado := soma.SomTudo(slice1, slice2, slice3)
		esperado := map[string]int{
			"[1 2]": 3,
			"[3 4]": 7,
			"[5 6]": 11,
		}

		for chave, valorEsperado := range esperado {
			if resultado[chave] != valorEsperado {
				t.Errorf("para a chave %s, resultado %d, esperado %d", chave, resultado[chave], valorEsperado)
			}
		}
	})

		t.Run("soma de múltiplos slices com um vazio", func(t *testing.T) {
		slice1 := []int{1, 2}
		slice2 := []int{3, 4}
		slice3 := []int{}
		resultado := soma.SomTudo(slice1, slice2, slice3)
		esperado := map[string]int{
			"[1 2]": 3,
			"[3 4]": 7,
			"[]": 0,
		}

		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado %v, esperado %v", resultado, esperado)
		}

		// for chave, valorEsperado := range esperado {
		// 	if resultado[chave] != valorEsperado {
		// 		t.Errorf("para a chave %s, resultado %d, esperado %d", chave, resultado[chave], valorEsperado)
		// 	}
		// }
	})
}
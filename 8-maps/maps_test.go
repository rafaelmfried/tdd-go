/*
	Aqui vemos algumas coisas interessantes sobre maps em Go:
	- Maps são tipos de referência, então quando passamos um map para uma função ou método,
	  estamos passando uma referência para o mapa original. Isso significa que alterações feitas
	  no mapa dentro da função ou método afetarão o mapa original.
*/

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
		dicionario := maps.NewDicionario()
		dicionario.Adicionar("teste", "isso é um teste")
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

func TestAdicionar(t *testing.T) {
	t.Run("Adiciona um novo par chave-valor ao dicionário", func(t *testing.T) {
		dicionario := maps.NewDicionario()
		dicionario.Adicionar("teste", "valor de teste")

		resultado, err := dicionario.Busca("teste")
		erroInexistente(t, err)

		esperado := "valor de teste"
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("Adicionar uma segunda vez uma chave existente retorna um erro", func(t *testing.T) {
		dicionario := maps.NewDicionario()
		_ = dicionario.Adicionar("teste", "valor de teste")
		err := dicionario.Adicionar("teste", "valor de teste")
		comparaErro(t, err, maps.ErroChaveExistente)
	})
}
package ponteiros_test

import (
	ponteiros "tdd/7-ponteiros"
	"testing"
)

func TestCarteira(t *testing.T) {
		t.Run("Depositar dinheiro na carteira", func(t *testing.T) {
				carteira := ponteiros.Carteira{}
				carteira.Depositar(10)

				resultado := carteira.Saldo()
				esperado := 10
				if resultado != esperado {
						t.Errorf("resultado %d, esperado %d", resultado, esperado)
				}
		})
	}
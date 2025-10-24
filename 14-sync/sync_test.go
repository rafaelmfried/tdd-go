package sync_test

import (
	sync "tdd/14-sync"
	"testing"
)

func TestContador(t *testing.T) {
	t.Run("incrementar o contador 3 vezes deve resultar em 3", func(t *testing.T) {
		contador := sync.NovoContador()

		contador.Incrementar()
		contador.Incrementar()
		contador.Incrementar()

		resultado := contador.Valor()
		valorEsperado := 3

		if resultado != valorEsperado {
			t.Errorf("resultado %d, valor esperado %d", resultado, valorEsperado)
		}
	})
}

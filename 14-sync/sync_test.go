package sync_test

import (
	"sync"
	. "tdd/14-sync"
	"testing"
)

func TestContador(t *testing.T) {
	t.Run("incrementar o contador 3 vezes deve resultar em 3", func(t *testing.T) {
		contador := NovoContador()

		contador.Incrementar()
		contador.Incrementar()
		contador.Incrementar()
		verificarContador(t, contador, 3)
	})

	t.Run("roda concorrentement em seguranca", func(t *testing.T) {
		contador := NovoContador()
		const totalDeRoutines = 100
		const incrementosPorRoutine = 1000
		// Maneira simples de esperar todas as goroutines terminarem
		var wg sync.WaitGroup
		wg.Add(totalDeRoutines)

		for i := 0; i < totalDeRoutines; i++ {
			go func() {
				for j := 0; j < incrementosPorRoutine; j++ {
					contador.Incrementar()
				}
				wg.Done()
			}()
		}

		wg.Wait()
		verificarContador(t, contador, totalDeRoutines*incrementosPorRoutine)
	})
}

func verificarContador(t *testing.T, contador *Contador, valorEsperado int) {
	t.Helper()
	resultado := contador.Valor()

	if resultado != valorEsperado {
		t.Errorf("resultado %d, valor esperado %d", resultado, valorEsperado)
	}
}

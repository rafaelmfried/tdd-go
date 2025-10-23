package helloworld_test

import (
	helloworld "tdd/2-hello-world"
	"testing"
)

func TestOla(t *testing.T) {
	t.Run("Diz olá para alguém", func(t *testing.T) {
		resultado := helloworld.Ola("Mundo", "português")
		esperado := "Olá, Mundo"

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("Diz olá para outra pessoa", func(t *testing.T) {
		resultado := helloworld.Ola("Rafael", "português")
		esperado := "Olá, Rafael"

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("Diz olá para ninguém", func(t *testing.T) {
		resultado := helloworld.Ola("", "português")
		esperado := "Olá, mundo"

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("Diz olá em português", func(t *testing.T) {
		resultado := helloworld.Ola("Gopher", "português")
		esperado := "Olá, Gopher"

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("Diz olá em espanhol", func(t *testing.T) {
		resultado := helloworld.Ola("Gopher", "espanhol")
		esperado := "Hola, Gopher"

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("Diz olá em inglês", func(t *testing.T) {
		resultado := helloworld.Ola("Gopher", "inglês")
		esperado := "Hello, Gopher"

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("Diz olá em idioma desconhecido", func(t *testing.T) {
		resultado := helloworld.Ola("Gopher", "francês")
		esperado := "Olá, Gopher"

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})
}

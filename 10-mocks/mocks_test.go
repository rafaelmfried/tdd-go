package mocks_test

import (
	"bytes"
	. "tdd/10-mocks"
	"testing"
)

type SleeperSpy struct {
	Calls int
}

func (s *SleeperSpy) Sleep() {
	s.Calls++
}

func TestMock(t *testing.T) {
	t.Run("Deve retornar a contagem", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SleeperSpy{}

		Contagem(buffer, sleeper)

		resultado := buffer.String()
		expected := "3\n2\n1\nVai!"

		if resultado != expected {
			t.Errorf("expected %s and result %s", expected, resultado)
		}

		if sleeper.Calls != 4 {
			t.Errorf("esperava 4 chamadas de sleep, mas recebeu %d", sleeper.Calls)
		}
	})
}
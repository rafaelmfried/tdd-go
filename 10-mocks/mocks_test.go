package mocks_test

import (
	"bytes"
	. "tdd/10-mocks"
	"testing"
)

func TestMock(t *testing.T) {
	t.Run("Deve retornar a contagem", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Contagem(buffer)

		resultado := buffer.String()
		expected := "3\n2\n1\nVai!"

		if resultado != expected {
			t.Errorf("expected %s and result %s", expected, resultado)
		}
	})
}
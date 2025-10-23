package iteracao_test

import (
	iteracao "tdd/iteracao"
	"testing"
)

func TestRepetir(t *testing.T) {
	repeticoes := iteracao.Repetir("a", 5)
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s', mas obteve '%s'", esperado, repeticoes)
	}
}
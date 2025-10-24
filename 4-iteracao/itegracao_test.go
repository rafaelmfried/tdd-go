package iteracao_test

import (
	iteracao "tdd/4-iteracao"
	"testing"
)

func TestRepetir(t *testing.T) {
	repeticoes := iteracao.Repetir("a", 5)
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s', mas obteve '%s'", esperado, repeticoes)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteracao.Repetir("a", 5)
	}
}
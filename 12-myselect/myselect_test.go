package myselect_test

import (
	myselect "tdd/12-myselect"
	"testing"
)

func TestCorredor(t *testing.T) {
	t.Run("teste de concorrencia com select", func(t *testing.T) {
		URLLenta := "http://www.facebook.com"
		URLRapida := "http://www.quii.co.uk"

		esperado := URLRapida

		obtido := myselect.Corredor(URLLenta, URLRapida)

		if obtido != esperado {
			t.Errorf("obtido: %q, esperado: %q", obtido, esperado)
		}
	})
}
package dependencias_test

import (
	"bytes"
	. "tdd/9-dependencias"
	"testing"
)

func TestDependencias(t *testing.T) {
		t.Run("Deve retornar frase padrao", func(t *testing.T) {
				expected := FrasePadrao
				var buffer bytes.Buffer
				Cumprimenta(&buffer, "")
				result := buffer.String()
				if result != expected {
					t.Errorf("Esperado %s, mas obteve %s", expected, result)
				}
		})

		t.Run("Deve cumprimentar o nome fornecido", func(t *testing.T) {
				nome := "Rafael"
				expected := "Olá, Rafael!"
				var buffer bytes.Buffer
				Cumprimenta(&buffer, nome)
				result := buffer.String()
				if result != expected {
					t.Errorf("Esperado %s, mas obteve %s", expected, result)
				}
		})

		t.Run("Deve receber um ponteiro", func(t *testing.T) {
			buffer := bytes.Buffer{}
			Cumprimenta(&buffer, "Chris")

			result := buffer.String()
			expected := "Olá, Chris!"

			if result != expected {
				t.Errorf("Esperado %s, mas obteve %s", expected, result)
			}
		})
	}
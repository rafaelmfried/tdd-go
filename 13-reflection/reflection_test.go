package reflection_test

import (
	"reflect"
	"testing"
)

func percorre(x interface{}, fn func(string)) {
	valor := reflect.ValueOf(x)
	campo := valor.Field(0)
	fn(campo.String())
}

func TestReflection(t *testing.T) {
	t.Run("placeholder", func (t *testing.T) {
    esperado := "Chris"

    var resultado []string
    x := struct {
        Nome string
    }{esperado}

    percorre(x, func(entrada string) {
        resultado = append(resultado, entrada)
    })

    if len(resultado) != 1 {
        t.Errorf("número incorreto de chamadas de função: resultado %d, esperado %d", len(resultado), 1)
    }

		if resultado[0] != esperado {
			t.Errorf("resultado incorreto, obtido: %q, esperado: %q", resultado[0], esperado)
		}
	})
}
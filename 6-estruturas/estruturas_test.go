package estruturas_test

import (
	estruturas "tdd/6-estruturas"
	"testing"
)

func TestPerimetro(t *testing.T) {
		t.Run("Retangulo", func(t *testing.T) {
			largura := 10.0
			altura := 15.0
			retangulo := estruturas.NewRetangulo(largura, altura)
			perimetro := estruturas.Perimetro(retangulo)
			esperado := 50.0

			if perimetro != esperado {
				t.Errorf("O perímetro calculado foi %f, mas o esperado é %f", perimetro, esperado)
			}
		})

		t.Run("Circulo", func(t *testing.T) {
			circulo := estruturas.NewCirculo(10.0)
			perimetro := estruturas.Perimetro(circulo)
			esperado := 62.83185307179586 // 2 * π * 10

			if perimetro != esperado {
				t.Errorf("O perímetro calculado foi %f, mas o esperado é %f", perimetro, esperado)
			}
		})
}

func TestArea(t *testing.T) {
		t.Run("Retangulo", func(t *testing.T) {
			largura := 10.0
			altura := 15.0
			retangulo := estruturas.NewRetangulo(largura, altura)
			area := retangulo.Area()
			esperado := 150.0

			if area != esperado {
				t.Errorf("A área calculada foi %f, mas a esperada é %f", area, esperado)
			}
		})

		t.Run("Circulo", func(t *testing.T) {
			circulo := estruturas.NewCirculo(10.0)
			area := circulo.Area()
			esperado := 314.1592653589793 // π * 10^2

			if area != esperado {
				t.Errorf("A área calculada foi %f, mas a esperada é %f", area, esperado)
			}
		})
}
package estruturas_test

import (
	estruturas "tdd/6-estruturas"
	"testing"
)

func TestPerimetro(t *testing.T) {
		testesPerimetro := []struct {
			forma estruturas.Forma
			esperado float64
		}{
			{estruturas.NewRetangulo(10.0, 15.0), 50.0},
			{estruturas.NewCirculo(10.0), 62.83185307179586},
			{estruturas.NewTriangulo(10.0, 15.0, 20.0), 45.0},
		}

		for _, tt := range testesPerimetro {
			perimetro := estruturas.Perimetro(tt.forma)
			t.Logf("Perímetro calculado para o %s: %f e o esperado: %f", tt.forma.GetNome(), perimetro, tt.esperado)
			if perimetro != tt.esperado {
				t.Errorf("O perímetro calculado para o %s foi %f, mas o esperado é %f", tt.forma.GetNome(), perimetro, tt.esperado)
			}
		}
}

func TestArea(t *testing.T) {
		testesArea := []struct {
			forma estruturas.Forma
			esperado float64
		}{
			{estruturas.NewRetangulo(10.0, 15.0), 150.0},
			{estruturas.NewCirculo(10.0), 314.1592653589793},
			{estruturas.NewTriangulo(3.0, 4.0, 5.0), 6.0},
		}

		for _, tt := range testesArea {
			area := estruturas.Area(tt.forma)
			t.Logf("Área calculada para o %s: %f e o esperado: %f", tt.forma.GetNome(), area, tt.esperado)
			if area != tt.esperado {
				t.Errorf("A área calculada para o %s foi %f, mas a esperada é %f", tt.forma.GetNome(), area, tt.esperado)
			}
		}
}
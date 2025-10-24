package estruturas

// Definimos a interface Forma que todas as formas geométricas devem implementar
type Forma interface {
	Perimetro() float64
	Area() float64
}

func Perimetro(forma Forma) float64 {
	return forma.Perimetro()
}

func Area(forma Forma) float64 {
	return forma.Area()
}

// Definimos a estrutura Retangulo
type Retangulo struct {
	largura float64
	altura  float64
}

func NewRetangulo(largura, altura float64) Retangulo {
	return Retangulo{largura: largura, altura: altura}
}

func (r Retangulo) Perimetro() float64 {
	return 2 * (r.largura + r.altura)
}

func (r Retangulo) Area() float64 {
	return r.largura * r.altura
}

// Definimos a estrutura Circulo
type Circulo struct {
	raio float64
}

func NewCirculo(raio float64) Circulo {
	return Circulo{raio: raio}
}

func (c Circulo) Perimetro() float64 {
	return 2 * 3.141592653589793 * c.raio
}

func (c Circulo) Area() float64 {
	return 3.141592653589793 * c.raio * c.raio
}
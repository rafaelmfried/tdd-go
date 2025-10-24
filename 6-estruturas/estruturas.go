package estruturas

// Definimos a interface Forma que todas as formas geométricas devem implementar
type Forma interface {
	Perimetro() float64
	Area() float64
	GetNome() string
}

const (
	retanguloTipo = "Retangulo"
	circuloTipo   = "Circulo"
	trianguloTipo = "Triangulo"
)

func Perimetro(forma Forma) float64 {
	return forma.Perimetro()
}

func Area(forma Forma) float64 {
	return forma.Area()
}

// Definimos a estrutura Retangulo
type Retangulo struct {
	nome    string
	largura float64
	altura  float64
}

func NewRetangulo(largura, altura float64) Retangulo {
	return Retangulo{nome: retanguloTipo, largura: largura, altura: altura}
}

func (r Retangulo) GetNome() string {
	return r.nome
}

func (r Retangulo) Perimetro() float64 {
	return 2 * (r.largura + r.altura)
}

func (r Retangulo) Area() float64 {
	return r.largura * r.altura
}

// Definimos a estrutura Circulo
type Circulo struct {
	nome string
	raio float64
}

func NewCirculo(raio float64) Circulo {
	return Circulo{nome: circuloTipo, raio: raio}
}

func (c Circulo) GetNome() string {
	return c.nome
}

func (c Circulo) Perimetro() float64 {
	return 2 * 3.141592653589793 * c.raio
}

func (c Circulo) Area() float64 {
	return 3.141592653589793 * c.raio * c.raio
}

// Definimos um triangulo como exemplo adicional
type Triangulo struct {
	nome  string
	ladoA float64
	ladoB float64
	ladoC float64
}

func NewTriangulo(ladoA, ladoB, ladoC float64) Triangulo {
	return Triangulo{nome: trianguloTipo,ladoA: ladoA, ladoB: ladoB, ladoC: ladoC}
}

func (t Triangulo) GetNome() string {
	return t.nome
}

func (t Triangulo) Perimetro() float64 {
	return t.ladoA + t.ladoB + t.ladoC
}

func (t Triangulo) Area() float64 {
	// Usando a fórmula de Herão para calcular a área do triângulo
	s := t.Perimetro() / 2
	return sqrt(s * (s - t.ladoA) * (s - t.ladoB) * (s - t.ladoC))
}

func sqrt(x float64) float64 {
	z := x
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}
package helloworld



const (
	idiomaPortugues = "português"
	idiomaIngles    = "inglês"
	idiomaEspanhol   = "espanhol"
)

const prefixos = map[string]string{
	idiomaPortugues: "Olá, ",
	idiomaIngles:    "Hello, ",
	idiomaEspanhol:  "Hola, ",
}

// const ( 
// 	prefixoOlaPortugues = "Olá, "
// 	prefixoOlaIngles    = "Hello, "
// 	prefixoOlaEspanhol  = "Hola, "
// )

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}

	switch idioma {
		case idiomaIngles:
			return prefixoOlaIngles + nome
		case idiomaEspanhol:
			return prefixoOlaEspanhol + nome
		case idiomaPortugues:
			return prefixoOlaPortugues + nome
		default:
			return prefixoOlaPortugues + nome
	}
}
package helloworld



const (
	idiomaPortugues = "português"
	idiomaIngles    = "inglês"
	idiomaEspanhol   = "espanhol"
)

var prefixos = map[string]string{
	idiomaPortugues: "Olá, ",
	idiomaIngles:    "Hello, ",
	idiomaEspanhol:  "Hola, ",
}


func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}

	prefixo, existe := prefixos[idioma]

	if !existe {
		prefixo = prefixos[idiomaPortugues]
	}

	return prefixo + nome
}
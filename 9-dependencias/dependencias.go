/*
	O IO Writer e o buffer bytes.Buffer implementam a interface io.Writer
*/

package dependencias

import (
	"fmt"
	io "io"
)

var FrasePadrao = "Olá, Mundo!"

func Cumprimenta(writter io.Writer, nome string) {
	if nome == "" {
		nome = "Mundo"
	}

	fmt.Fprintf(writter, "Olá, %s!", nome)
}
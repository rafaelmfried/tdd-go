package mocks

import (
	"fmt"
	"io"
	"time"
)

const inicioContagem = 3
const fraseFinal = "Vai!"

func Contagem(saida io.Writer) {
	for i := inicioContagem; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(saida, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(saida, fraseFinal)
}
package mocks

import (
	"fmt"
	"io"
	"time"
)

const inicioContagem = 3
const fraseFinal = "Vai!"

type Sleeper interface {
	Sleep()
}

type SleeperPadrao struct{}

func (s *SleeperPadrao) Sleep() {
	time.Sleep(1 * time.Second)
}

func Contagem(saida io.Writer, sleeper Sleeper) {
    for i := inicioContagem; i > 0; i-- {
			sleeper.Sleep()
			fmt.Fprintln(saida, i)
    }

    sleeper.Sleep()
    fmt.Fprint(saida, fraseFinal)
}

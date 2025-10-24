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
type SleeperConfiguravel struct {
	duration time.Duration
	sleep func(time.Duration)
}

func (s *SleeperConfiguravel) Sleep() {
	s.sleep(s.duration)
}

func NewSleeper(duration time.Duration, sleeper func(time.Duration)) *SleeperConfiguravel {
	return &SleeperConfiguravel{duration, sleeper}
}

func Contagem(saida io.Writer, sleeper Sleeper) {
    for i := inicioContagem; i > 0; i-- {
			sleeper.Sleep()
			fmt.Fprintln(saida, i)
    }

    sleeper.Sleep()
    fmt.Fprint(saida, fraseFinal)
}

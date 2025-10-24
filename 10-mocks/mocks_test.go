package mocks_test

import (
	"bytes"
	"reflect"
	. "tdd/10-mocks"
	mocks "tdd/10-mocks"
	"testing"
	"time"
)

type SleeperSpy struct {
	Calls int
}

func (s *SleeperSpy) Sleep() {
	s.Calls++
}

type SpyContagemOperacoes struct {
	Operacoes []string
}

func (s *SpyContagemOperacoes) Sleep() {
	s.Operacoes = append(s.Operacoes, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Operacoes = append(s.Operacoes, escrita)
	return len(p), nil
}

const (
	escrita = "escrita"
	pausa   = "sleep"
)
type TempoSpy struct {
	durationSlept time.Duration
}

func (t *TempoSpy) Sleep(duration time.Duration) {
	t.durationSlept += duration
}

func TestMock(t *testing.T) {
	t.Run("Deve retornar a contagem", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SleeperSpy{}

		Contagem(buffer, sleeper)

		resultado := buffer.String()
		expected := "3\n2\n1\nVai!"

		if resultado != expected {
			t.Errorf("expected %s and result %s", expected, resultado)
		}

		if sleeper.Calls != 4 {
			t.Errorf("esperava 4 chamadas de sleep, mas recebeu %d", sleeper.Calls)
		}
	})

	t.Run("Deve garantir a ordem correta entre escrita e pausa", func(t *testing.T) {
		spy := &SpyContagemOperacoes{}

		Contagem(spy, spy)

		expected := []string{
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
		}

		if !reflect.DeepEqual(expected, spy.Operacoes) {
			t.Errorf("esperava-se a sequência %v, mas recebeu %v", expected, spy.Operacoes)
		}
	})

	t.Run("Deve configurar o sleeper com o tempo correto", func(t *testing.T) {
		sleepTime := 5 * time.Second

		tempoSpy := &TempoSpy{}
		sleeper := mocks.NewSleeper(sleepTime, tempoSpy.Sleep)

		sleeper.Sleep()

		if tempoSpy.durationSlept != sleepTime {
			t.Errorf("esperava-se dormir por %v, mas dormiu por %v", sleepTime, tempoSpy.durationSlept)
		}
	})


}
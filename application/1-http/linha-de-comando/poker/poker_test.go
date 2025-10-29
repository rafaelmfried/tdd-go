package poker_test

import (
	"bytes"
	"fmt"
	"strings"
	"tdd/application/1-http/linha-de-comando/cli"
	"tdd/application/1-http/linha-de-comando/helpers"
	"testing"
	"time"
)

type scheduledAlert struct {
	at time.Duration
	amount int
}

func (s *scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{duration, amount})
}
func TestPoker(t *testing.T) {
	t.Run("deve agendar a impressao dos valores dos blinds para 5 jogadores", func(t *testing.T) {
		in := strings.NewReader("Rafael venceu\n")
		stdout := &bytes.Buffer{}
		armazenamento := &helpers.EsbocoArmazenamentoJogador{}
		dummySpyAlerter := &SpyBlindAlerter{}

		cli := cli.NovoCLI(armazenamento, in, stdout, dummySpyAlerter)
		cli.JogarPoquer()

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprintf("%d scheduled for %v", want.amount, want.at), func(t *testing.T) {
				if len(dummySpyAlerter.alerts) <= i {
					t.Fatalf("alerta %d nao agendado %v", i, dummySpyAlerter.alerts)
				}

				got := dummySpyAlerter.alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})

	t.Run("deve pedir o numero de jogadores", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		cli := cli.NovoCLI(
			&helpers.EsbocoArmazenamentoJogador{},
			&bytes.Buffer{},
			stdout,
			&SpyBlindAlerter{},
		)

		cli.JogarPoquer()

		got := stdout.String()
		want := "Please enter the number of players: "

		if got != want {
			t.Errorf("esperava '%s', mas recebeu '%s'", want, got)
		}
	})

}

func assertScheduledAlert(t *testing.T, got, want scheduledAlert) {
	t.Helper()
	if got.amount != want.amount {
		t.Errorf("esperava o valor do alerta ser %d, mas recebeu %d", want.amount, got.amount)
	}

	if got.at != want.at {
		t.Errorf("esperava o tempo agendado do alerta ser %v, mas recebeu %v", want.at, got.at)
	}
}

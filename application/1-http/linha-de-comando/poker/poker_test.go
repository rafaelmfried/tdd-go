package poker_test

import (
	"bytes"
	"fmt"
	"tdd/application/1-http/linha-de-comando/cli"
	"tdd/application/1-http/linha-de-comando/helpers"
	"tdd/application/1-http/linha-de-comando/poker"
	"testing"
	"time"
)

const PlayerPrompt = cli.PlayerPrompt

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
		armazenamento := &helpers.EsbocoArmazenamentoJogador{}
		dummySpyAlerter := &SpyBlindAlerter{}
		game := poker.NewGame(dummySpyAlerter, armazenamento)


		game.Start(5)

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

		checkSchedulingCases(cases, t, dummySpyAlerter)
	})

	t.Run("deve pedir o numero de jogadores", func(t *testing.T) {
		in := helpers.UserSends("5", "Rafael venceu")
		stdout := &bytes.Buffer{}
		game := poker.NewGame(&SpyBlindAlerter{}, &helpers.EsbocoArmazenamentoJogador{})
		cli := cli.NovoCLI(
			in,
			stdout,
			game,
		)

		cli.JogarPoquer()

		got := stdout.String()
		want := PlayerPrompt

		if got != want {
			t.Errorf("esperava '%s', mas recebeu '%s'", want, got)
		}
	})

	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
        blindAlerter := &SpyBlindAlerter{}
				dummyPlayerStore := &helpers.EsbocoArmazenamentoJogador{}
        game := poker.NewGame(blindAlerter, dummyPlayerStore)

        game.Start(7)

        cases := []scheduledAlert{
            {at: 0 * time.Second, amount: 100},
            {at: 12 * time.Minute, amount: 200},
            {at: 24 * time.Minute, amount: 300},
            {at: 36 * time.Minute, amount: 400},
        }

        checkSchedulingCases(cases, t, blindAlerter)
    })

		t.Run("o jogo deve terminar", func(t *testing.T) {
			store := &helpers.EsbocoArmazenamentoJogador{}
			dummyAlerter := &SpyBlindAlerter{}
			game := poker.NewGame(dummyAlerter, store)
			winner := "Rafael"

			game.Finish(winner)

			assertPlayerWin(t, store, winner)
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

func checkSchedulingCases(cases []scheduledAlert, t *testing.T, dummySpyAlerter *SpyBlindAlerter) {
			for i, want := range cases {
			t.Run(fmt.Sprintf("%d scheduled for %v", want.amount, want.at), func(t *testing.T) {
				if len(dummySpyAlerter.alerts) <= i {
					t.Fatalf("alerta %d nao agendado %v", i, dummySpyAlerter.alerts)
				}

				got := dummySpyAlerter.alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	}

	func assertPlayerWin(t *testing.T, store *helpers.EsbocoArmazenamentoJogador, winner string) {
		t.Helper()

		if len(store.RegistrosVitorias) != 1 {
			t.Fatalf("esperava 1 vitoria registrada, mas recebeu %d", len(store.RegistrosVitorias))
		}

		if store.RegistrosVitorias[0] != winner {
			t.Errorf("esperava o vencedor ser '%s', mas recebeu '%s'", winner, store.RegistrosVitorias[0])
		}
	}

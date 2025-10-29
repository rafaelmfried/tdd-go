package poker

import (
	"fmt"
	"strings"
	"tdd/application/1-http/linha-de-comando/cli"
	"tdd/application/1-http/linha-de-comando/helpers"
	"testing"
	"time"
)
func TestPoker(t *testing.T) {
	t.Run("deve agendar a impressao dos valores dos blinds para 5 jogadores", func(t *testing.T) {
		in := strings.NewReader("Rafael venceu\n")
		armazenamento := &helpers.EsbocoArmazenamentoJogador{}
		dummySpyAlerter := &SpyBlindAlerter{}
		
		cli := cli.NovoCLI(armazenamento, in, dummySpyAlerter)
		cli.JogarPoquer()

		cases := []struct{
			expectedScheduleTime time.Duration
			expectedAmount int
		}{
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

		for i, c := range cases {
			t.Run(fmt.Sprintf("%d scheduled for %v", c.expectedAmount, c.expectedScheduleTime), func(t *testing.T) {
				if len(dummySpyAlerter.alerts) <= i {
					t.Fatalf("alerta %d nao agendado", i)
				}

				alert := dummySpyAlerter.alerts[i]

				amountGot := alert.amount
				if amountGot != c.expectedAmount {
					t.Errorf("esperava o valor do alerta %d ser %d, mas recebeu %d", i, c.expectedAmount, amountGot)
				}

				scheduledTimeGot := alert.scheduleAt
				if scheduledTimeGot != c.expectedScheduleTime {
					t.Errorf("esperava o tempo agendado do alerta %d ser %v, mas recebeu %v", i, c.expectedScheduleTime, scheduledTimeGot)
				}
			})
		}
	})
}

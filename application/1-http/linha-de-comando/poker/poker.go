package poker

import "time"

type SpyBlindAlerter struct {
	alerts []struct{
		scheduleAt time.Duration
		amount int
	}
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, struct{
		scheduleAt time.Duration
		amount int
	}{duration, amount})
}
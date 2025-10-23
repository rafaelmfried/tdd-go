package iteracao

import "strings"

func Repetir(target string, times int) string {
	var builder strings.Builder

	builder.Grow(len(target) * times)
	for i := 0; i < times; i++ {
		builder.WriteString(target)
	}

	return builder.String()
}
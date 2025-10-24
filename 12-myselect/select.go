package myselect

import (
	"net/http"
	"time"
)

func Corredor(url1, url2 string) (vencedor string) {
	inicioA := time.Now()
	http.Get(url1)
	duracaoA := time.Since(inicioA)

	inicioB := time.Now()
	http.Get(url2)
	duracaoB := time.Since(inicioB)

	if duracaoA < duracaoB {
		return url1
	}
	return url2
}
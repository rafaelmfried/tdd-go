package myselect

import (
	"net/http"
)

func Corredor(url1, url2 string) (vencedor string) {
	// duracaoA := medirTempoRequisicao(url1)
	// duracaoB := medirTempoRequisicao(url2)

	// if duracaoA < duracaoB {
	// 	return url1
	// }
	// return url2
	select {
	case <-ping(url1):
		return url1
	case <-ping(url2):
		return url2
	}
}

func ping(URL string) chan bool {
	ch := make(chan bool)

	go func() {
		http.Get(URL)
		ch <- true
	}()

	return ch
}

// func medirTempoRequisicao(url string) time.Duration {
// 	inicio := time.Now()
// 	http.Get(url)
// 	return time.Since(inicio)
// }
package myselect

import (
	"fmt"
	"net/http"
	"time"
)

var ErrTimeout = fmt.Errorf("tempo de espera excedido na requisicao")
var responseTimeout = 2 * time.Second

func Corredor(url1, url2 string) (vencedor string, erro error) {
	return CorredorConfiguravel(url1, url2, responseTimeout)
}

func CorredorConfiguravel(url1, url2 string, timeLimit time.Duration) (vencedor string, erro error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeLimit):
		return "", ErrTimeout
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
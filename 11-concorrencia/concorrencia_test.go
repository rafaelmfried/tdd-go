/*
* Copyright © 2024 Rafael Friederick <rafaelfriederick@gmail.com>
* go test -bench=. -benchmem ./11-concorrencia
* Agora falando em concorrencia em go para rodar processos no bloqueantes de forma concorrente
* precisamos usar goroutines e para isso colocamos go na frente da chamada da funcao
* go test -race ajuda a identificar condicoes de concorrencia que podem causar problemas
* por conta disso que usamos canais para sincronizar a escrita no mapa resultados
 */

package concorrencia_test

import (
	"reflect"
	concorrencia "tdd/11-concorrencia"
	"testing"
	"time"
)

func mockVerificadorWebsite(url string) bool {
	return url == "http://blog.golang.org" || url == "http://golang.org"
}

func slowVerificadorWebsite(url string) bool {
	time.Sleep(2 * time.Second)
	return url == "http://blog.golang.org" || url == "http://golang.org"
}

func TestVerificaWebsites(t *testing.T) {
	websites := []string{
		"http://blog.golang.org",
		"http://golang.org",
		"http://invalid-url.org",
	}

	resultadosEsperados := map[string]bool{
		"http://blog.golang.org":   true,
		"http://golang.org":        true,
		"http://invalid-url.org":   false,
	}

	resultado := concorrencia.VerificaWebsites(mockVerificadorWebsite, websites)

	if reflect.DeepEqual(resultado, resultadosEsperados) {
		t.Logf("Resultados coincidem")
	} else {
		t.Errorf("Resultados não coincidem: esperado %v, mas obteve %v", resultadosEsperados, resultado)
	}
}

func BeanchmarkVerificaWebsites(b *testing.B) {
	b.Run("Beanch VerificaWebsites com mockVerificadorWebsite", func(b *testing.B) {
		websites := []string{
			"http://blog.golang.org",
			"http://golang.org",
			"http://invalid-url.org",
		}

		for i := 0; i < b.N; i++ {
			concorrencia.VerificaWebsites(mockVerificadorWebsite, websites)
		}
	})

	b.Run("VerificaWebsites com slowVerificadorWebsite", func(b *testing.B) {
		websites := []string{
			"http://blog.golang.org",
			"http://golang.org",
			"http://invalid-url.org",
		}

		for i := 0; i < b.N; i++ {
			concorrencia.VerificaWebsites(slowVerificadorWebsite, websites)
		}
	})
}
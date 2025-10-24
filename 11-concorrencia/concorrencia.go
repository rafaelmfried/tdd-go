/*
Aqui criamos os canais entre os processo pai a funcao VerificaWebsites e os filhos que sao as goroutines
*/
package concorrencia

type VerificadorWebsite func(string) bool

type resultado struct {
		string
	 	bool
}

func VerificaWebsites(verificador VerificadorWebsite, websites []string) map[string]bool {
	resultados := make(map[string]bool)
	resultChannel := make(chan resultado)

	for _, url := range websites {
		go func(u string) {
			resultChannel <- resultado{u, verificador(u)}
		}(url)
	}

	for i := 0; i < len(websites); i++ {
		resultado := <-resultChannel
		resultados[resultado.string] = resultado.bool
	}

	return resultados
}
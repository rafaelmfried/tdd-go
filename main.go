package main

import (
	"net/http"
	dependencias "tdd/9-dependencias"
)

func HandlerCumprimenta(w http.ResponseWriter, r *http.Request) {
	dependencias.Cumprimenta(w, "Rafael")
}

func main() {
	// armazenamento := server.NovoArmazenamentoJogadorDoArquivo()
	// server := server.NewServidorJogador(armazenamento)

	// if err := http.ListenAndServe(":2312", server); err != nil {
	// 	log.Fatalf("binding port 2312 %v", err)
	// }
}
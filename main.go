package main

import (
	"log"
	"net/http"
	dependencias "tdd/9-dependencias"
	server "tdd/application/1-http"
)

func HandlerCumprimenta(w http.ResponseWriter, r *http.Request) {
	dependencias.Cumprimenta(w, "Rafael")
}

func main() {
	armazenamento := server.NovoArmazenamentoJogadorInMemory()
	server := server.NewServidorJogador(armazenamento)

	if err := http.ListenAndServe(":2312", server); err != nil {
		log.Fatalf("binding port 2312 %v", err)
	}
}
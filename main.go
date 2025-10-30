package main

import (
	"log"
	"net/http"
	"os"
	dependencias "tdd/9-dependencias"
	"tdd/application/1-http/linha-de-comando/armazenamento"
	"tdd/application/1-http/linha-de-comando/servidor"
)

func HandlerCumprimenta(w http.ResponseWriter, r *http.Request) {
	dependencias.Cumprimenta(w, "Rafael")
}

func main() {
	db, err := os.OpenFile("jogo.db.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("opening db file %v", err)
	}
	defer db.Close()
	store, err := armazenamento.NovoArmazenamentoJogadorDoArquivo(db)
	if err != nil {
		log.Fatalf("creating player storage %v", err)
	}
	server := servidor.NewServidorJogador(store)

	if err := http.ListenAndServe(":2312", server); err != nil {
		log.Fatalf("binding port 2312 %v", err)
	}
}

package main

import (
	"log"
	"net/http"
	"os"
	armazenamento "tdd/application/1-http/linha-de-comando/armazenamento"
	servidor "tdd/application/1-http/linha-de-comando/servidor"
)

const fileName = "jogo.db.json"

func main() {
	db, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("falha ao abrir %s %v", fileName, err)
	}

	armazenamento, err := armazenamento.NovoArmazenamentoJogadorDoArquivo(db)

	if err != nil {
		log.Fatalf("falha ao criar armazenamento do jogador a partir do arquivo %s %v", fileName, err)
	}

	servidor := servidor.NewServidorJogador(armazenamento)

	if err := http.ListenAndServe(":3212", servidor); err != nil {
		log.Fatalf("nao foi possivel escutar na porta 3212 %v", err)
	}
}
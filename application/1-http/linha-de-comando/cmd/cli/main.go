package main

import (
	"fmt"
	"log"
	"os"
	"tdd/application/1-http/linha-de-comando/armazenamento"
	"tdd/application/1-http/linha-de-comando/cli"
	"tdd/application/1-http/linha-de-comando/poker"
)

const fileName = "jogo.db.json"

func main() {
	fmt.Println("Vamos jogar poquer")
	fmt.Println("Digite {Nome} venceu para registrar uma vitoria")

	db, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}

	armazenamento, err := armazenamento.NovoArmazenamentoJogadorDoArquivo(db)

	if err != nil {
		log.Fatalf("Erro ao criar o armazenamento do jogador: %v", err)
	}

	game := poker.NewGame(poker.BlindAlerterFunc(poker.StdOutAlerter), armazenamento)

	cli := cli.NovoCLI(os.Stdin, os.Stdout, game)
	cli.JogarPoquer()
}

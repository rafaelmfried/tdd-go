package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	mocks "tdd/10-mocks"
	dependencias "tdd/9-dependencias"
	server "tdd/application/1-http"
	"time"
)

func HandlerCumprimenta(w http.ResponseWriter, r *http.Request) {
	dependencias.Cumprimenta(w, "Rafael")
}

func newLine() {
	fmt.Println()
}

func main() {
	// err := http.ListenAndServe(":3456", http.HandlerFunc(HandlerCumprimenta))

	// if err != nil {
	// 	panic(err)
	// }

	dependencias.Cumprimenta(os.Stdout, "Rafael")
	newLine()
	// sleeper := &mocks.SleeperPadrao{}
	duration := 1 * time.Second
	sleeper := mocks.NewSleeper(duration, time.Sleep)
	mocks.Contagem(os.Stdout, sleeper)
	newLine()

	server := &server.ServidorJogador{}

	if err := http.ListenAndServe(":2312", server); err != nil {
		log.Fatalf("binding port 2312 %v", err)
	}
}
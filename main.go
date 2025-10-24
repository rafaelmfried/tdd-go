package main

import (
	"fmt"
	"net/http"
	"os"
	mocks "tdd/10-mocks"
	dependencias "tdd/9-dependencias"
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
	sleeper := &mocks.SleeperPadrao{}
	mocks.Contagem(os.Stdout, sleeper)
	newLine()
}
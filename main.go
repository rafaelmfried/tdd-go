package main

import (
	"net/http"
	"os"
	dependencias "tdd/9-dependencias"
)

func HandlerCumprimenta(w http.ResponseWriter, r *http.Request) {
	dependencias.Cumprimenta(w, "Rafael")
}

func main() {
	err := http.ListenAndServe(":3456", http.HandlerFunc(HandlerCumprimenta))

	if err != nil {
		panic(err)
	}

	dependencias.Cumprimenta(os.Stdout, "Rafael")
}
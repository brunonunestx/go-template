package main

import (
	"log"
	"net/http"

	apiRouter "gateway/internal/http"
)

func main() {
	router := apiRouter.NewRouter()

	log.Println("Servidor rodando em :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

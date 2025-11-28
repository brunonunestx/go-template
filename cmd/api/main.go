package main

import (
	"log"
	"net/http"

	apiRouter "template/internal/http"
)

func main() {
	router := apiRouter.NewRouter()

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"

	"log"
	"net/http"
)

func main() {
	config.Carregar()
	fmt.Println("Iniciando o servidor na porta", config.Porta)
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}

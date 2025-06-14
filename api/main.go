package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"

	"log"
	"net/http"
)

/*
//Funcao inicial para criar um chave usada para definir chave no .env
func init() {
	chave := make([]byte, 64)

	//aqui gera numeros aleatorios para chave
	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	//usar essa chave para o segredo do token
	fmt.Println(stringBase64)

}*/

func main() {
	config.Carregar()
	fmt.Println("Iniciando o servidor na porta", config.Porta)
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}

//PARA EXECUTAR TEM QUE ESTAR COM O TERMINAL DENTRO DA PASTA DA API E RODAR COMANDO
//go run main.go

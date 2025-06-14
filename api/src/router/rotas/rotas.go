package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI                string `json:"uri"`
	Metodo             string `json:"metodo"`
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios //Essa variavel faz a ligação com o usuarios.go
	rotas = append(rotas, rotaLogin)
	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}

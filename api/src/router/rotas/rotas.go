package rotas

import (
	"api/src/middlewares"
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
	rotas = append(rotas, rotasPublicacoes...)
	for _, rota := range rotas {
		if rota.RequerAutenticacao {
			r.HandleFunc(
				rota.URI,
				middlewares.Cors(middlewares.Logger(middlewares.Autenticar(rota.Funcao))),
			).Methods(rota.Metodo, "OPTIONS")
		} else {
			r.HandleFunc(
				rota.URI,
				middlewares.Cors(middlewares.Logger(rota.Funcao)),
			).Methods(rota.Metodo)
		}
	}

	return r
}

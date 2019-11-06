package aplicacao

import "github.com/gorilla/mux"

// NewRotuter - Retorna todas as rotas(apis) da aplicacao.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	return router
}

package visao

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter - Retorna todas as rotas(apis) da aplicacao.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range rotas {

		//HANDLER ATUH
		var authHandler http.Handler
		authHandler = route.HandlerFunc
		authHandler = basicAuth(authHandler, route.Name)

		if route.Name == "PostLoginUsuario" {
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Handler(route.HandlerFunc).
				Name(route.Name)
			continue
		}

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc).
			Handler(authHandler)
	}
	return router
}

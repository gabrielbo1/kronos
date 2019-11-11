package visao

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter - Retorna todas as rotas(apis) da aplicacao.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range rotas {

		//HANDLER DE LOG
		var logHandler http.Handler
		logHandler = route.HandlerFunc
		logHandler = logger(logHandler, route.Name)

		if route.Name == "Index" ||
			route.Name == "PostLoginUsuario" {
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(logHandler)
			continue
		}

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(logHandler)
	}
	return router
}

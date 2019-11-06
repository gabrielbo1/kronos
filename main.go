package main

import (
	"log"
	"net/http"

	"github.com/gabrielbo1/kronos/aplicacao"
	"github.com/gabrielbo1/kronos/infraestrutura"
	"github.com/gabrielbo1/kronos/infraestrutura/repositorio"
	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	if infraestrutura.Config.DiretorioScripts != "" {
		repositorio.ShcemaUpdate(infraestrutura.Config.DiretorioScripts)
	}

	staticDir := "./static"
	stripPrefix := http.StripPrefix("/", http.FileServer(http.Dir(staticDir)))

	router := aplicacao.NewRouter()
	router.PathPrefix("/").Handler(handlers.CompressHandler(stripPrefix))
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":80", cors.AllowAll().Handler(router)))
}

package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/gabrielbo1/kronos/infraestrutura"
	"github.com/gabrielbo1/kronos/infraestrutura/repositorio"
	"github.com/gabrielbo1/kronos/visao"
	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func init() {
	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func main() {
	if infraestrutura.Config.DiretorioScripts != "" {
		repositorio.ShcemaUpdate(infraestrutura.Config.DiretorioScripts)
	}

	staticDir := "./static"
	stripPrefix := http.StripPrefix("/", http.FileServer(http.Dir(staticDir)))

	router := visao.NewRouter()
	router.PathPrefix("/").Handler(handlers.CompressHandler(stripPrefix))
	http.Handle("/", router)

	port := os.Getenv("PORT") // Heroku provides the port to bind to
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, cors.AllowAll().Handler(router)))
}

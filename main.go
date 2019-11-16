package main

import (
	"net/http"
	"os"

	"github.com/gabrielbo1/kronos/infraestrutura/repositorio"
	"github.com/rs/cors"
	"github.com/urfave/negroni"

	log "github.com/sirupsen/logrus"

	"github.com/gabrielbo1/kronos/infraestrutura"
	"github.com/gabrielbo1/kronos/visao"
	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
)

func init() {
	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	infraestrutura.ConfigInit()
	repositorio.ShcemaUpdate(infraestrutura.Config.DiretorioScripts)

	staticDir := "./static"
	stripPrefix := http.StripPrefix("/", http.FileServer(http.Dir(staticDir)))

	router := visao.NewRouter()
	router.PathPrefix("/").Handler(handlers.CompressHandler(stripPrefix))
	http.Handle("/", router)

	port := os.Getenv("PORT") // Heroku provides the port to bind to
	if port == "" {
		port = "8080"
	}

	n := negroni.Classic()
	n.UseHandler(cors.AllowAll().Handler(router))
	n.Run(":" + port)
}

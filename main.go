package main

import (
	"github.com/gabrielbo1/kronos/infraestrutura"
	"github.com/gabrielbo1/kronos/infraestrutura/repositorio"
	_ "github.com/lib/pq"
)

func main() {
	if infraestrutura.Config.DiretorioScripts != "" {
		repositorio.ShcemaUpdate(infraestrutura.Config.DiretorioScripts)
	}
}

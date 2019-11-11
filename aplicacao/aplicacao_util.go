package aplicacao

import (
	"log"

	"github.com/gabrielbo1/kronos/dominio"
)

// TrataErroConexao Em caso de erro na conecção com o banco de dados
// para a aplicação e lança log.
func TrataErroConexao(erroApp *dominio.Erro, err error) *dominio.Erro {
	if erroApp != nil {
		return erroApp
	}
	log.Println("\tErro transação banco de dados \t " + err.Error())
	return &dominio.Erro{Codigo: "SERVICO_UTIL_10", Mensagem: "Erro transação banco de dados", Err: err}
}

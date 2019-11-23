package aplicacao

import (
	"database/sql"

	"github.com/gabrielbo1/kronos/dominio"
	"github.com/gabrielbo1/kronos/infraestrutura/repositorio"
)

var repPonto = repositorio.NewPontoRepositorio()

// CadastrarPonto - Cadastrar novo ponto, registrar ponto.
func CadastrarPonto(ponto *dominio.Ponto) (errDominio *dominio.Erro) {
	if errDominio = dominio.NewPonto(ponto); errDominio != nil {
		return errDominio
	}

	if errTX := repositorio.Transact(repositorio.DB, func(tx *sql.Tx) error {
		ponto.ID, errDominio = repPonto.Save(tx, *ponto)
		return dominio.OnError(errDominio)
	}); errTX != nil {
		return TrataErroConexao(errDominio, errTX)
	}
	return errDominio
}

// AtualizarPonto - Atualizar ponto.
func AtualizarPonto(ponto *dominio.Ponto) (errDominio *dominio.Erro) {
	if errDominio = dominio.NewPonto(ponto); errDominio != nil {
		return errDominio
	}

	if errTX := repositorio.Transact(repositorio.DB, func(tx *sql.Tx) error {
		errDominio = repPonto.Update(tx, *ponto)
		return dominio.OnError(errDominio)
	}); errTX != nil {
		return TrataErroConexao(errDominio, errTX)
	}
	return errDominio
}

// ApagarPonto - Apagar ponto.
func ApagarPonto(ponto *dominio.Ponto) (errDominio *dominio.Erro) {
	if errTX := repositorio.Transact(repositorio.DB, func(tx *sql.Tx) error {
		errDominio = repPonto.Delete(tx, *ponto)
		return dominio.OnError(errDominio)
	}); errTX != nil {
		return TrataErroConexao(errDominio, errTX)
	}
	return errDominio
}

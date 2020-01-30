package aplicacao

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/now"
	"time"

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

// BuscarPontoDia - Busca os pontos de um usuario registrado em um determinado dia.
func BuscarPontoDia(idUsuario int, data time.Time) (pontos []dominio.Ponto, errDominio *dominio.Erro) {
	if errTX := repositorio.Transact(repositorio.DB, func(tx *sql.Tx) error {
		pontos, errDominio = repPonto.FindByData(tx, idUsuario, data)
		return dominio.OnError(errDominio)
	}); errTX != nil {
		return pontos, TrataErroConexao(errDominio, errTX)
	}

	for i := range pontos {
		if dt, err := now.Parse(pontos[i].Data); err == nil {
			pontos[i].Data = fmt.Sprintf("%2d/%2d/%d %2d:%2d", dt.Day(), dt.Month(), dt.Year(), dt.Hour(), dt.Minute())
		}

	}

	return pontos, errDominio
}

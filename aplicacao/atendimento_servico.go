package aplicacao

import (
	"database/sql"

	"github.com/gabrielbo1/kronos/dominio"
	"github.com/gabrielbo1/kronos/infraestrutura/repositorio"
)

var repAtendimento = repositorio.NewRepositorioAtendimento()

// CadastrarAtendimento - Cadastra novo atendimento.
func CadastrarAtendimento(atendimento *dominio.Atendimento) (errDominio *dominio.Erro) {
	if errDominio = dominio.NewAtendimento(atendimento); errDominio != nil {
		return errDominio
	}

	if errTX := repositorio.Transact(repositorio.DB, func(tx *sql.Tx) error {
		atendimento.ID, errDominio = repAtendimento.Save(tx, *atendimento)
		return nil
	}); errTX != nil {
		return TrataErroConexao(errDominio, errTX)
	}

	return errDominio
}

// AtualizarAtendimento - Atualizar atendimento.
func AtualizarAtendimento(atendimento *dominio.Atendimento) (errDominio *dominio.Erro) {
	if errDominio = dominio.NewAtendimento(atendimento); errDominio != nil {
		return errDominio
	}

	if errTX := repositorio.Transact(repositorio.DB, func(tx *sql.Tx) error {
		errDominio = repAtendimento.Update(tx, *atendimento)
		return nil
	}); errTX != nil {
		return TrataErroConexao(errDominio, errTX)
	}

	return errDominio
}

// ApagarAtendimento - Apagar atendimento.
func ApagarAtendimento(atendimento *dominio.Atendimento) (errDominio *dominio.Erro) {
	if errTX := repositorio.Transact(repositorio.DB, func(tx *sql.Tx) error {
		errDominio = repAtendimento.Delete(tx, *atendimento)
		return nil
	}); errTX != nil {
		return TrataErroConexao(errDominio, errTX)
	}

	return errDominio
}

// BuscarAtendimentoIdUsuario - Busca todos atendimentos associados ao usuario.
func BuscarAtendimentoIdUsuario(idUsuario int) (atendimentos []dominio.Atendimento, errDominio *dominio.Erro) {
	if errTX := repositorio.Transact(repositorio.DB, func(tx *sql.Tx) error {
		atendimentos, errDominio = repAtendimento.FindByIdUsuario(tx, idUsuario)
		return nil
	}); errTX != nil {
		return atendimentos, TrataErroConexao(errDominio, errTX)
	}

	return atendimentos, errDominio
}

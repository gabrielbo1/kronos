package aplicacao

import (
	"database/sql"

	"github.com/gabrielbo1/kronos/dominio"
	"github.com/gabrielbo1/kronos/infraestrutura/repositorio"
)

var repEmpresa = repositorio.NewEmpresaRepositorio()

// CadastrarEmpresa - Cadastra nova empresa.
func CadastrarEmpresa(empresa *dominio.Empresa) (errDominio *dominio.Erro) {
	if errDominio = dominio.NewEmpresa(empresa); errDominio != nil {
		return errDominio
	}

	if errTX := repositorio.Transact(repositorio.DB, func(tx *sql.Tx) error {
		empresa.ID, errDominio = repEmpresa.Save(tx, *empresa)
		return dominio.OnError(errDominio)
	}); errTX != nil {
		return TrataErroConexao(errDominio, errTX)
	}

	return errDominio
}

// AtualizarEmpresa - Atualiza empresa.
func AtualizarEmpresa(empresa *dominio.Empresa) (errDominio *dominio.Erro) {
	if errDominio = dominio.NewEmpresa(empresa); errDominio != nil {
		return errDominio
	}

	if errTX := repositorio.Transact(repositorio.DB, func(tx *sql.Tx) error {
		errDominio = repEmpresa.Update(tx, *empresa)
		return dominio.OnError(errDominio)
	}); errTX != nil {
		return TrataErroConexao(errDominio, errTX)
	}

	return errDominio
}

// ApagarEmpresa - Apagar empresa.
func ApagarEmpresa(empresa *dominio.Empresa) (errDominio *dominio.Erro) {
	if errTX := repositorio.Transact(repositorio.DB, func(tx *sql.Tx) error {
		errDominio = repEmpresa.Delete(tx, *empresa)
		return dominio.OnError(errDominio)
	}); errTX != nil {
		return TrataErroConexao(errDominio, errTX)
	}

	return errDominio
}

// BuscaEmpresas - Busca todas empresas cadastradas.
func BuscaEmpresas() (empresas []dominio.Empresa, errDominio *dominio.Erro) {
	if errTX := repositorio.Transact(repositorio.DB, func(tx *sql.Tx) error {
		empresas, errDominio = repEmpresa.FindAll(tx)
		return dominio.OnError(errDominio)
	}); errTX != nil {
		return empresas, TrataErroConexao(errDominio, errTX)
	}

	return empresas, errDominio
}

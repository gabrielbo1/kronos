package repositorio

import (
	"database/sql"

	"github.com/gabrielbo1/kronos/dominio"
)

// EmpresaRepPostgres - Tipo Repositorio Empresa.
type empresaRepPostgres string

// Save - Salva uma nova empresa, implementacao Postgresql.
func (empresaRepPostgres) Save(tx *sql.Tx, entidade dominio.Empresa) (int, *dominio.Erro) {
	return 0, nil
}

// Update - Atualiza um empresa, implementacao Postgresql.
func (empresaRepPostgres) Update(tx *sql.Tx, entidade dominio.Empresa) *dominio.Erro {
	return nil
}

// Delete - Deleta uma empresa, implementacao Postgresql.
func (empresaRepPostgres) Delete(tx *sql.Tx, entidade dominio.Empresa) *dominio.Erro {
	return nil
}

// FindAll - Busca todas empresas cadastradas, implementacao Postgresql.
func (empresaRepPostgres) FindAll(tx *sql.Tx) (entidades []dominio.Empresa, erro *dominio.Erro) {
	return nil, nil
}

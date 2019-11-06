package repositorio

import (
	"database/sql"

	"github.com/gabrielbo1/kronos/dominio"
)

// EmpresaRepPostgres - Tipo Repositorio Empresa.
type empresaRepPostgres string

// Save - Salva uma nova empresa, implementacao Postgresql.
func (empresaRepPostgres) Save(tx *sql.Tx, entidade dominio.Empresa) (id int, errDomin *dominio.Erro) {
	sqlInsert := `INSERT INTO empresa(id, nome_empresa, ativa) VALUES ($1, $2, $3) RETURNING ID`
	stmt, err := prepararStmt(ctx, tx, "empresaRepPostgres", "Save", sqlInsert)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	if ok, errDomin := scanParamStmt("empresaRepPostgres", "Save", stmt, func(stmt *sql.Stmt) error {
		return stmt.QueryRowContext(ctx, &entidade.ID, &entidade.Nome, &entidade.Ativa).Scan(&entidade.ID)
	}); !ok {
		return 0, errDomin
	}
	return id, nil
}

// Update - Atualiza um empresa, implementacao Postgresql.
func (empresaRepPostgres) Update(tx *sql.Tx, entidade dominio.Empresa) *dominio.Erro {
	sqlUptade := "UPDATE empresa SET id=$1, nome_empresa=$2, ativa=$3 WHERE id=$4;"
	stmt, err := prepararStmt(ctx, tx, "empresaRepPostgres", "Update", sqlUptade)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if ok, errDomain := scanParamStmt("Update", "Save", stmt, func(stmt *sql.Stmt) error {
		_, err := stmt.ExecContext(ctx, &entidade.ID, entidade.Nome, entidade.Ativa, entidade.ID)
		return err
	}); !ok {
		return errDomain
	}
	return nil
}

// Delete - Deleta uma empresa, implementacao Postgresql.
func (empresaRepPostgres) Delete(tx *sql.Tx, entidade dominio.Empresa) *dominio.Erro {
	sqlDelete := "DELETE FROM empresa WHERE id=$1"
	stmt, err := prepararStmt(ctx, tx, "empresaRepPostgres", "Delete", sqlDelete)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if ok, errDomain := scanParamStmt("empresaRepPostgres", "Delete", stmt, func(stmt *sql.Stmt) error {
		_, err := stmt.ExecContext(ctx, entidade.ID)
		return err
	}); !ok {
		return errDomain
	}
	return nil
}

// FindAll - Busca todas empresas cadastradas, implementacao Postgresql.
func (empresaRepPostgres) FindAll(tx *sql.Tx) (entidades []dominio.Empresa, erro *dominio.Erro) {
	return nil, nil
}

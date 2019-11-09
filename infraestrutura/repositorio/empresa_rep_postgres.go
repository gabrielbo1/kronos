package repositorio

import (
	"database/sql"

	"github.com/gabrielbo1/kronos/dominio"
)

// EmpresaRepPostgres - Tipo Repositorio Empresa.
type empresaRepPostgres string

// Save - Salva uma nova empresa, implementacao Postgresql.
func (empresaRepPostgres) Save(tx *sql.Tx, entidade dominio.Empresa) *dominio.Erro {
	sqlInsert := `INSERT INTO empresa(nome_empresa, ativa) VALUES ($1, $2) RETURNING ID`
	stmt, err := prepararStmt(ctx, tx, "empresaRepPostgres", "Save", sqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if ok, errDomin := scanParamStmt("empresaRepPostgres", "Save", stmt, func(stmt *sql.Stmt) error {
		return stmt.QueryRowContext(ctx, &entidade.Nome, &entidade.Ativa).Scan(&entidade.ID)
	}); !ok {
		return errDomin
	}
	return nil
}

// Update - Atualiza um empresa, implementacao Postgresql.
func (empresaRepPostgres) Update(tx *sql.Tx, entidade dominio.Empresa) *dominio.Erro {
	sqlUptade := "UPDATE empresa SET id=$1, nome_empresa=$2, ativa=$3 WHERE id=$4"
	stmt, err := prepararStmt(ctx, tx, "empresaRepPostgres", "Update", sqlUptade)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if ok, errDomain := scanParamStmt("empresaRepPostgres", "Update", stmt, func(stmt *sql.Stmt) error {
		_, err := stmt.ExecContext(ctx, &entidade.ID, entidade.Nome, entidade.Ativa, entidade.ID)
		return err
	}); !ok {
		return errDomain
	}
	return nil
}

// Delete - Deleta uma empresa, implementacao Postgresql.
func (empresaRepPostgres) Delete(tx *sql.Tx, entidade dominio.Empresa) *dominio.Erro {
	return delete(tx, "empresaRepPostgres", "empresa", entidade.ID)
}

// FindAll - Busca todas empresas cadastradas, implementacao Postgresql.
func (empresaRepPostgres) FindAll(tx *sql.Tx) (entidades []dominio.Empresa, erro *dominio.Erro) {
	rows, errTx := tx.QueryContext(ctx, "SELECT id, nome_empresa, ativa FROM empresa")
	if errTx != nil {
		return []dominio.Empresa{}, &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro FindAll empresaRepPostgres função FindAll", Err: errTx}
	}
	empresas, err := parseEmpresaPostgres(rows)
	if err != nil {
		return []dominio.Empresa{}, &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro FindAll empresaRepPostgres função FindAll", Err: errTx}
	}
	return empresas, nil
}

func parseEmpresaPostgres(rows *sql.Rows) ([]dominio.Empresa, error) {
	defer rows.Close()
	var results []dominio.Empresa
	for rows.Next() {
		emp := dominio.Empresa{}
		if err := rows.Scan(&emp.ID, &emp.Nome, &emp.Ativa); err != nil {
			return nil, err
		}
		results = append(results, emp)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return results, nil
}

package repositorio

import (
	"database/sql"

	"github.com/gabrielbo1/kronos/dominio"
)

// rotinaRepPostgres - Tipo Repositorio Rotina para Postgresql.
type rotinaRepPostgres string

func (rotinaRepPostgres) Save(tx *sql.Tx, entidade dominio.Rotina) (int, *dominio.Erro) {
	sqlInsert := `INSERT INTO rotina(rotina) VALUES ($1) RETURNING ID`
	stmt, err := prepararStmt(ctx, tx, "rotinaRepPostgres", "Save", sqlInsert)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	id := 0
	if ok, errDomin := scanParamStmt("rotinaRepPostgres", "Save", stmt, func(stmt *sql.Stmt) error {
		return stmt.QueryRowContext(ctx, &entidade.Rotina).Scan(&id)
	}); !ok {
		return 0, errDomin
	}
	return id, nil
}

func (rotinaRepPostgres) Update(tx *sql.Tx, entidade dominio.Rotina) *dominio.Erro {
	sqlUptade := "UPDATE rotina SET rotina=$1 WHERE id=$2;"
	stmt, err := prepararStmt(ctx, tx, "rotinaRepPostgres", "Update", sqlUptade)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if ok, errDomain := scanParamStmt("rotinaRepPostgres", "Update", stmt, func(stmt *sql.Stmt) error {
		_, err := stmt.ExecContext(ctx, &entidade.Rotina, entidade.ID)
		return err
	}); !ok {
		return errDomain
	}
	return nil
}

func (rotinaRepPostgres) Delete(tx *sql.Tx, entidade dominio.Rotina) *dominio.Erro {
	return delete(tx, "rotinaRepPostgres", "rotina", entidade.ID)
}

func (rotinaRepPostgres) FindAll(tx *sql.Tx) (entidades []dominio.Rotina, erro *dominio.Erro) {
	rows, errTx := tx.QueryContext(ctx, "SELECT id, rotina FROM rotina")
	if errTx != nil {
		return []dominio.Rotina{}, &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro FindAll rotinaRepPostgres função FindAll", Err: errTx}
	}
	empresas, err := parseRotinaPostgres(rows)
	if err != nil {
		return []dominio.Rotina{}, &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro FindAll rotinaRepPostgres função FindAll", Err: errTx}
	}
	return empresas, nil
}

func parseRotinaPostgres(rows *sql.Rows) ([]dominio.Rotina, error) {
	defer rows.Close()
	var results []dominio.Rotina
	for rows.Next() {
		rot := dominio.Rotina{}
		if err := rows.Scan(&rot.ID, &rot.Rotina); err != nil {
			return nil, err
		}
		results = append(results, rot)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return results, nil
}

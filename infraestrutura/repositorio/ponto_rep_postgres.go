package repositorio

import (
	"database/sql"

	"github.com/gabrielbo1/kronos/dominio"
)

// PontoRepPostgres - Tipo Repositorio Ponto.
type pontoRepPostgres string

// Save - Salva uma nova ponto, implementacao Postgresql.
func (pontoRepPostgres) Save(tx *sql.Tx, entidade dominio.Ponto) (int, *dominio.Erro) {
	sqlInsert := "INSERT INTO ponto(idusuario, data) VALUES ($1, $2)"
	stmt, err := prepararStmt(ctx, tx, "pontoRepPostgres", "Save", sqlInsert)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	id := 0
	if ok, errDomin := scanParamStmt("pontoRepPostgres", "Save", stmt, func(stmt *sql.Stmt) error {
		return stmt.QueryRowContext(ctx, &entidade.Usuario.ID, ajustarDataPostgres(entidade.Data)).Scan(&id)
	}); !ok {
		return 0, errDomin
	}
	return id, nil
}

// Update - Atualiza um ponto, implementacao Postgresql.
func (pontoRepPostgres) Update(tx *sql.Tx, entidade dominio.Ponto) *dominio.Erro {
	sqlUptade := "UPDATE ponto SET idusuario=$1, data=$2 WHERE id=$3"
	stmt, err := prepararStmt(ctx, tx, "pontoRepPostgres", "Update", sqlUptade)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if ok, errDomain := scanParamStmt("pontoRepPostgres", "Update", stmt, func(stmt *sql.Stmt) error {
		_, err := stmt.ExecContext(ctx, &entidade.Usuario.ID, ajustarDataPostgres(entidade.Data), entidade.ID)
		return err
	}); !ok {
		return errDomain
	}
	return nil
}

// Delete - Deleta um ponto, implementacao Postgresql.
func (pontoRepPostgres) Delete(tx *sql.Tx, entidade dominio.Ponto) *dominio.Erro {
	return delete(tx, "pontoRepPostgres", "ponto", entidade.ID)
}

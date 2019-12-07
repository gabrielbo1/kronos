package repositorio

import (
	"database/sql"
	"time"

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

// FindByData - Buscar
func (pontoRepPostgres) FindByData(tx *sql.Tx, idUsuario int, data time.Time) ([]dominio.Ponto, *dominio.Erro) {
	sqlSelect := `SELECT id, idusuario, data FROM ponto 
			 WHERE idusuario = $1 AND date_part('day', data)   = $2 
								  AND date_part('month', data) = $3 
								  AND date_part('year', data)  = $4`
	stmt, err := prepararStmt(ctx, tx, "pontoRepPostgres", "FindByData", sqlSelect)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	var rows *sql.Rows
	var err1 error
	if ok, errDomin := scanParamStmt("pontoRepPostgres", "FindByData", stmt, func(stmt *sql.Stmt) error {
		if rows, err1 = stmt.QueryContext(ctx, &idUsuario, data.Day(), data.Month(), data.Year()); err1 != nil {
			return err1
		}
		return nil
	}); !ok {
		return nil, errDomin
	}

	result, err1 := parsePonto(rows)
	if err1 != nil {
		return nil, &dominio.Erro{Codigo: "SQLUTIL_REP20",
			Mensagem: "Erro Atendimento pontoRepPostgres função FindByData",
			Err:      err1}
	}

	return result, nil
}

func parsePonto(rows *sql.Rows) ([]dominio.Ponto, error) {
	defer rows.Close()
	var result []dominio.Ponto
	for rows.Next() {
		p := dominio.Ponto{}
		p.Usuario = dominio.Usuario{}
		if err := rows.Scan(&p.ID, &p.Usuario.ID, &p.Data); err != nil {
			return nil, err
		}
		result = append(result, p)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return result, nil
}

package repositorio

import (
	"database/sql"
	"encoding/json"

	"github.com/gabrielbo1/kronos/dominio"
)

// rotinaRepPostgres - Tipo Repositorio Atendimento para Postgresql.
type atendimentoRepPostgres string

func (atendimentoRepPostgres) Save(tx *sql.Tx, entidade dominio.Atendimento) (int, *dominio.Erro) {
	sqlInsert := `INSERT INTO atendimento(idUsuario, idCliente, horarios_atendimento, status_atendimento, observacao) 
					VALUES ($1, $2, $3, $4, $5) RETURNING ID`
	stmt, err := prepararStmt(ctx, tx, "atendimentoRepPostgres", "Save", sqlInsert)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var jsonByte []byte
	var errJSON error
	if jsonByte, errJSON = json.Marshal(entidade.HorariosAtendimento); err != nil {
		return 0, &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro atendimentoRepPostgres função Save", Err: errJSON}
	}

	id := 0
	if ok, errDomin := scanParamStmt("atendimentoRepPostgres", "Save", stmt, func(stmt *sql.Stmt) error {
		return stmt.QueryRowContext(ctx, &entidade.Usuario.ID, &entidade.Cliente.ID,
			string(jsonByte), &entidade.StatusAtendimento, &entidade.Observacao).Scan(&id)
	}); !ok {
		return 0, errDomin
	}
	return id, nil
}

func (atendimentoRepPostgres) Update(tx *sql.Tx, entidade dominio.Atendimento) *dominio.Erro {
	sqlUptade := `UPDATE atendimento SET idUsuario=$1, idCliente=$2, horarios_atendimento=$3, status_atendimento=$4, observacao=$5 
					WHERE id=$6`
	stmt, err := prepararStmt(ctx, tx, "atendimentoRepPostgres", "Update", sqlUptade)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var jsonByte []byte
	var errJSON error
	if jsonByte, errJSON = json.Marshal(entidade.HorariosAtendimento); err != nil {
		return &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro atendimentoRepPostgres função Save", Err: errJSON}
	}

	if ok, errDomain := scanParamStmt("atendimentoRepPostgres", "Update", stmt, func(stmt *sql.Stmt) error {
		_, err := stmt.ExecContext(ctx, entidade.Usuario.ID, &entidade.Cliente.ID,
			string(jsonByte), &entidade.StatusAtendimento, &entidade.Observacao, &entidade.ID)
		return err
	}); !ok {
		return errDomain
	}
	return nil
}

func (atendimentoRepPostgres) Delete(tx *sql.Tx, entidade dominio.Atendimento) *dominio.Erro {
	return delete(tx, "atendimentoRepPostgres", "atendimento", entidade.ID)
}

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

func (atendimentoRepPostgres) FindByIdUsuario(tx *sql.Tx, id int) (entidades []dominio.Atendimento, erro *dominio.Erro) {
	sqlQuery := `select id, idusuario, idcliente, 
       					horarios_atendimento, status_atendimento, observacao 
						from atendimento where idusuario = $1`
	stmt, errDomin := prepararStmt(ctx, tx, "atendimentoRepPostgres", "FindByIdUsuario", sqlQuery)
	if errDomin != nil {
		return []dominio.Atendimento{}, errDomin
	}
	defer stmt.Close()

	rows, errTx := stmt.QueryContext(ctx, id)
	if errTx != nil {
		return []dominio.Atendimento{}, &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro Atendimento atendimentoRepPostgres função FindByIdUsuario", Err: errTx}
	}

	entidades, err := parseAtendimentoPostgres(rows)
	if err != nil {
		return []dominio.Atendimento{}, &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro Login usuarioRepPostgres função Login", Err: err}
	}
	return entidades, nil
}

func parseAtendimentoPostgres(rows *sql.Rows) ([]dominio.Atendimento, error) {
	defer rows.Close()
	var result []dominio.Atendimento
	for rows.Next() {
		atn := dominio.Atendimento{}
		atn.Usuario = dominio.Usuario{}
		atn.Cliente = dominio.Empresa{}
		atn.HorariosAtendimento = []dominio.Intervalo{}
		jsonHorariosAtendimento := ""

		if err := rows.Scan(&atn.ID, &atn.Usuario.ID, &atn.Cliente.ID,
			&jsonHorariosAtendimento, &atn.StatusAtendimento, &atn.Observacao); err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(jsonHorariosAtendimento), &atn.HorariosAtendimento); err != nil {
			return nil, err
		}

		result = append(result, atn)
	}
	return result, nil
}

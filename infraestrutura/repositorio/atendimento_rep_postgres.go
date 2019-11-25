package repositorio

import (
	"database/sql"
	"log"
	"time"

	"github.com/gabrielbo1/kronos/dominio"
)

// rotinaRepPostgres - Tipo Repositorio Atendimento para Postgresql.
type atendimentoRepPostgres string

func (atendimentoRepPostgres) Save(tx *sql.Tx, entidade dominio.Atendimento) (int, *dominio.Erro) {
	sqlInsert := `INSERT INTO atendimento(idUsuario, idCliente, status_atendimento, observacao) 
					VALUES ($1, $2, $3, $4) RETURNING ID`
	stmt, err := prepararStmt(ctx, tx, "atendimentoRepPostgres", "Save", sqlInsert)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	id := 0
	if ok, errDomin := scanParamStmt("atendimentoRepPostgres", "Save", stmt, func(stmt *sql.Stmt) error {
		return stmt.QueryRowContext(ctx, &entidade.Usuario.ID,
			&entidade.Cliente.ID,
			&entidade.StatusAtendimento,
			&entidade.Observacao).Scan(&id)
	}); !ok {
		return 0, errDomin
	}

	if errDomin := saveIntervalo(tx, id, entidade.HorariosAtendimento); errDomin != nil {
		return 0, errDomin
	}

	return id, nil
}

func (atendimentoRepPostgres) Update(tx *sql.Tx, entidade dominio.Atendimento) *dominio.Erro {
	sqlUptade := `UPDATE atendimento SET idUsuario=$1, idCliente=$2, status_atendimento=$3, observacao=$4 
					WHERE id=$5`
	stmt, err := prepararStmt(ctx, tx, "atendimentoRepPostgres", "Update", sqlUptade)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if ok, errDomain := scanParamStmt("atendimentoRepPostgres", "Update", stmt, func(stmt *sql.Stmt) error {
		_, err := stmt.ExecContext(ctx, entidade.Usuario.ID, &entidade.Cliente.ID,
			&entidade.StatusAtendimento, &entidade.Observacao, &entidade.ID)
		return err
	}); !ok {
		return errDomain
	}

	if errDomin := saveIntervalo(tx, entidade.ID, entidade.HorariosAtendimento); errDomin != nil {
		return errDomin
	}
	return nil
}

func (atendimentoRepPostgres) Delete(tx *sql.Tx, entidade dominio.Atendimento) *dominio.Erro {
	if err := deleteNameColumn(tx, "atendimentoRepPostgres", "INTERVALO", "IDATENDIMENTO", entidade.ID); err != nil {
		return err
	}
	return delete(tx, "atendimentoRepPostgres", "atendimento", entidade.ID)
}

func (atendimentoRepPostgres) FindByIdUsuario(tx *sql.Tx, id int) (entidades []dominio.Atendimento, erro *dominio.Erro) {
	sqlQuery := `select atendimento.id, idusuario, u.nome, idcliente, e.nome_empresa, 
       					horarios_atendimento, status_atendimento, observacao 
						from atendimento 
						inner join usuario u on u.id = idusuario
						inner join empresa e on e.id =  idcliente
						where idusuario = $1`
	stmt, errDomin := prepararStmt(ctx, tx, "atendimentoRepPostgres", "FindByIdUsuario", sqlQuery)
	if errDomin != nil {
		return []dominio.Atendimento{}, errDomin
	}
	defer stmt.Close()

	rows, errTx := stmt.QueryContext(ctx, id)
	if errTx != nil {
		return []dominio.Atendimento{}, &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro Atendimento atendimentoRepPostgres função FindByIdUsuario", Err: errTx}
	}

	entidades, err := parseAtendimentoPostgres(tx, rows)
	if err != nil {
		return []dominio.Atendimento{}, &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro Login usuarioRepPostgres função Login", Err: err}
	}
	return entidades, nil
}

func (atendimentoRepPostgres) FindByIdUsuarioPaginado(tx *sql.Tx, idUsuario int, paginaSolicitada dominio.Pagina) (pagina dominio.Pagina, errDominio *dominio.Erro) {
	var totalRegistro int = 0
	pagina.NumPagina = paginaSolicitada.NumPagina
	pagina.QtdPorPagina = paginaSolicitada.QtdPorPagina

	sqlCount := `select count(*) 
				 from atendimento a 
				 inner join usuario u on u.id = a.idusuario 
                 inner join empresa e on e.id = a.idcliente 
				 where a.idusuario = $1`

	stmt, err := tx.PrepareContext(ctx, sqlCount)
	if err != nil {
		log.Println(err)
		return paginaSolicitada, &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro Atendimento atendimentoRepPostgres função FindByIdUsuarioPaginado", Err: err}
	}

	defer stmt.Close()
	if err = stmt.QueryRowContext(ctx, idUsuario).Scan(&totalRegistro); err != nil {
		log.Println(err)
		return paginaSolicitada, &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro Atendimento atendimentoRepPostgres função FindByIdUsuarioPaginado", Err: err}
	}
	pagina.TotalRegistro = totalRegistro
	pagina.TotalPagina = dominio.CalcQtdPaginas(pagina.TotalRegistro, pagina.QtdPorPagina)

	sqlBusca := `select a.id, 
						a.idusuario, 
						u.nome, 
						a.idcliente, 
						e.nome_empresa,
						a.status_atendimento, 
						a.observacao 
				 from atendimento a 
				 inner join usuario u on u.id = a.idusuario 
                 inner join empresa e on e.id = a.idcliente 
				 where a.idusuario = $1 
				 ORDER BY a.id DESC `
	sqlBusca += geraSqlOfsset(pagina)

	stmtBusca, err := tx.PrepareContext(ctx, sqlBusca)
	if err != nil {
		log.Println(err)
		return paginaSolicitada, &dominio.Erro{Codigo: "SQLUTIL_REP20",
			Mensagem: "Erro Atendimento atendimentoRepPostgres função FindByIdUsuarioPaginado",
			Err:      err}
	}

	defer stmtBusca.Close()

	rows, errTx := stmtBusca.QueryContext(ctx, idUsuario)
	if errTx != nil {
		log.Println(err)
		return paginaSolicitada, &dominio.Erro{Codigo: "SQLUTIL_REP20",
			Mensagem: "Erro Atendimento atendimentoRepPostgres função FindByIdUsuarioPaginado",
			Err:      errTx}
	}

	pagina.Conteudo, err = parseAtendimentoPostgres(tx, rows)
	if err != nil {
		log.Println(err)
		return paginaSolicitada, &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro Atendimento atendimentoRepPostgres função FindByIdUsuarioPaginado", Err: err}
	}
	return pagina, nil
}

func (atendimentoRepPostgres) FindByIdUsuarioPaginadoLike(tx *sql.Tx, idUsuario int, like string, paginaSolicitada dominio.Pagina) (pagina dominio.Pagina, errDominio *dominio.Erro) {
	var totalRegistro int = 0
	pagina.NumPagina = paginaSolicitada.NumPagina
	pagina.QtdPorPagina = paginaSolicitada.QtdPorPagina

	sqlCount := `select count(*) 
				 from atendimento a 
				 inner join usuario u on u.id = a.idusuario 
                 inner join empresa e on e.id = a.idcliente 
				 where a.idusuario = $1 `
	sqlCount += " AND (" + likeAtendimento(like) + ") "

	stmt, err := tx.PrepareContext(ctx, sqlCount)
	if err != nil {
		log.Println(err)
		return paginaSolicitada, &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro Atendimento atendimentoRepPostgres função FindByIdUsuarioPaginado", Err: err}
	}

	defer stmt.Close()
	if err = stmt.QueryRowContext(ctx, idUsuario).Scan(&totalRegistro); err != nil {
		log.Println(err)
		return paginaSolicitada, &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro Atendimento atendimentoRepPostgres função FindByIdUsuarioPaginado", Err: err}
	}
	pagina.TotalRegistro = totalRegistro
	pagina.TotalPagina = dominio.CalcQtdPaginas(pagina.TotalRegistro, pagina.QtdPorPagina)

	sqlBusca := `select a.id, 
						a.idusuario, 
						u.nome, 
						a.idcliente, 
						e.nome_empresa,
						a.status_atendimento, 
						a.observacao 
				 from atendimento a 
				 inner join usuario u on u.id = a.idusuario 
                 inner join empresa e on e.id = a.idcliente 
				 where a.idusuario = $1 `

	sqlBusca += " AND (" + likeAtendimento(like) + ") "
	sqlBusca += " ORDER BY a.id DESC"

	sqlBusca += geraSqlOfsset(pagina)

	stmtBusca, err := tx.PrepareContext(ctx, sqlBusca)
	if err != nil {
		log.Println(err)
		return paginaSolicitada, &dominio.Erro{Codigo: "SQLUTIL_REP20",
			Mensagem: "Erro Atendimento atendimentoRepPostgres função FindByIdUsuarioPaginado",
			Err:      err}
	}

	defer stmtBusca.Close()

	rows, errTx := stmtBusca.QueryContext(ctx, idUsuario)
	if errTx != nil {
		log.Println(err)
		return paginaSolicitada, &dominio.Erro{Codigo: "SQLUTIL_REP20",
			Mensagem: "Erro Atendimento atendimentoRepPostgres função FindByIdUsuarioPaginado",
			Err:      errTx}
	}

	pagina.Conteudo, err = parseAtendimentoPostgres(tx, rows)
	if err != nil {
		log.Println(err)
		return paginaSolicitada, &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro Atendimento atendimentoRepPostgres função FindByIdUsuarioPaginado", Err: err}
	}
	return pagina, nil
}

func likeAtendimento(texto string) string {
	like := " UPPER(u.nome) LIKE UPPER('%" + texto + "%') OR"
	like += " UPPER(e.nome_empresa) LIKE UPPER('%" + texto + "%')"
	return like
}

func saveIntervalo(tx *sql.Tx, idAtendimento int, intervalos []dominio.Intervalo) *dominio.Erro {
	var errDomin *dominio.Erro
	for i := 0; i < len(intervalos) && errDomin == nil; i++ {
		if intervalos[i].ID == 0 {
			errDomin = inserIntervalo(tx, idAtendimento, intervalos[i])
			continue
		}
		errDomin = updateIntervalo(tx, intervalos[i])
	}
	return errDomin
}

func inserIntervalo(tx *sql.Tx, idAtendimento int, intervalo dominio.Intervalo) *dominio.Erro {
	sqlInsert := `INSERT INTO INTERVALO (IDATENDIMENTO, DATA_INICIO, DATA_FIM) VALUES ($1, $2, $3) RETURNING ID `
	stmt, err := prepararStmt(ctx, tx, "atendimentoRepPostgres", "inserIntervalo", sqlInsert)
	if err != nil {
		return err
	}

	defer stmt.Close()
	if ok, errDomin := scanParamStmt("atendimentoRepPostgres", "saveIntervalo", stmt, func(stmt *sql.Stmt) error {
		return stmt.QueryRowContext(ctx, &idAtendimento,
			ajustarDataPostgres(intervalo.DataInicio),
			ajustarDataPostgres(intervalo.DataFim)).Scan(&intervalo.ID)
	}); !ok {
		return errDomin
	}

	return nil
}

func updateIntervalo(tx *sql.Tx, intervalo dominio.Intervalo) *dominio.Erro {
	sqlUpdate := "UPDATE INTERVALO SET DATA_FIM=$1 WHERE ID = $2"
	stmt, err := prepararStmt(ctx, tx, "atendimentoRepPostgres", "updateIntervalo", sqlUpdate)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if ok, errDomin := scanParamStmt("atendimentoRepPostgres", "updateIntervalo", stmt, func(stmt *sql.Stmt) error {
		_, err := stmt.ExecContext(ctx, ajustarDataPostgres(intervalo.DataFim), &intervalo.ID)
		return err
	}); !ok {
		return errDomin
	}
	return nil
}

func findInervalosAtendimento(tx *sql.Tx, idAtendimento int) ([]dominio.Intervalo, error) {
	sqlSeletIntervalo := `SELECT ID, DATA_INICIO, DATA_FIM FROM INTERVALO WHERE IDATENDIMENTO = $1`
	stmt, err := prepararStmt(ctx, tx, "atendimentoRepPostgres", "findInervalosAtendimento", sqlSeletIntervalo)
	if err != nil {
		return nil, err.Err
	}
	defer stmt.Close()

	var rows *sql.Rows
	var err1 error
	if ok, errDomin := scanParamStmt("atendimentoRepPostgres", "findInervalosAtendimento", stmt, func(stmt *sql.Stmt) error {
		if rows, err1 = stmt.QueryContext(ctx, &idAtendimento); err1 != nil {
			return err1
		}
		return nil
	}); !ok {
		return nil, errDomin.Err
	}

	defer rows.Close()
	var result []dominio.Intervalo
	for rows.Next() {
		interv := dominio.Intervalo{}
		var dtInicio sql.NullTime
		var dtFim sql.NullTime
		if err2 := rows.Scan(&interv.ID, &dtInicio, &dtFim); err2 != nil {
			return nil, err2
		}

		if dtInicio.Valid {
			interv.DataInicio = dtInicio.Time.Format(time.RFC3339)
		}

		if dtFim.Valid {
			interv.DataFim = dtFim.Time.Format(time.RFC3339)
		}

		result = append(result, interv)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return result, nil
}

func parseAtendimentoPostgres(tx *sql.Tx, rows *sql.Rows) ([]dominio.Atendimento, error) {
	defer rows.Close()
	var result []dominio.Atendimento
	for rows.Next() {
		atn := dominio.Atendimento{}
		atn.Usuario = dominio.Usuario{}
		atn.Cliente = dominio.Empresa{}
		atn.HorariosAtendimento = []dominio.Intervalo{}

		if err := rows.Scan(&atn.ID,
			&atn.Usuario.ID,
			&atn.Usuario.Nome,
			&atn.Cliente.ID,
			&atn.Cliente.Nome,
			&atn.StatusAtendimento,
			&atn.Observacao); err != nil {
			return nil, err
		}
		result = append(result, atn)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	for i := range result {
		var err error
		if result[i].HorariosAtendimento, err = findInervalosAtendimento(tx, result[i].ID); err != nil {
			return nil, err
		}
	}

	return result, nil
}

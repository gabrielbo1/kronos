package repositorio

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/gabrielbo1/kronos/dominio"
)

//BooleanToString - Boolean para sql.
func BooleanToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// StringToString - String para sql.
func StringToString(s string) string {
	return "'" + s + "'"
}

//InCreate - Cria string sql para in.
func InCreate(ids []int) string {
	in := "("
	for i := 0; i < len(ids); i++ {
		in += strconv.Itoa(ids[i])
		in += ","
	}
	in = in[0 : len(in)-1]
	in += ")"
	return in
}

//AjustarDataPostgres - Ajusta data de acordo com o padrao para Postgresql.
func AjustarDataPostgres(dataString string) string {
	var data time.Time
	var err error
	if data, err = time.Parse(time.RFC3339, dataString); err != nil {
		return dataString
	}
	ano, mes, day := data.Date()
	return strconv.Itoa(ano) + "-" + strconv.Itoa(int(mes)) + "-" + strconv.Itoa(day)
}

// PrepararStmt - Funcao generica para preparar o Stmt.
func PrepararStmt(ctx context.Context, tx *sql.Tx, nomeRep, nomeFunc, query string) (*sql.Stmt, *dominio.Erro) {
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Println("SQLUTIL_REP10: Erro ao preparar consulta, repositório " + nomeRep + " função " + nomeFunc)
		log.Println(err)
		return nil, &dominio.Erro{"SQLUTIL_REP10", "Erro ao preparara consulta, repositório " + nomeRep + " função " + nomeFunc, err}
	}
	return stmt, nil
}

// ScanStmt - Funcao generica para realizar o parser dos campos da consulta feita com stmt.
func ScanStmt(ctx context.Context, nomeRep, nomeFunc string, stmt *sql.Stmt, args ...interface{}) (bool, *dominio.Erro) {
	err := stmt.QueryRowContext(ctx).Scan(args)
	switch {
	case err == sql.ErrNoRows:
		return false, nil
	case err != nil:
		log.Println("Erro realizar Scan Smmt, repositório " + nomeRep + " função " + nomeFunc)
		log.Println(err)
		return false, &dominio.Erro{"SQLUTIL_REP20", "Erro realizar Scan Smmt, repositório " + nomeRep + " função " + nomeFunc, err}
	default:
		return true, nil
	}
	return true, nil
}

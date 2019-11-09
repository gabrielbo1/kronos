package repositorio

import (
	"database/sql"
	"github.com/gabrielbo1/kronos/dominio"
)

// rotinaRepPostgres - Tipo Repositorio Atendimento para Postgresql.
type atendimentoRepPostgres string

func (atendimentoRepPostgres) Save(tx *sql.Tx, entidade dominio.Atendimento) (int, *dominio.Erro) {
	return 0, nil
}

func (atendimentoRepPostgres) Update(tx *sql.Tx, entidade dominio.Atendimento) *dominio.Erro {
	return nil
}

func (atendimentoRepPostgres) Delete(tx *sql.Tx, entidade dominio.Atendimento) *dominio.Erro {
	return nil
}

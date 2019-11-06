package repositorio

import (
	"database/sql"

	"github.com/gabrielbo1/kronos/dominio"
)

// PontoRepPostgres - Tipo Repositorio Ponto.
type pontoRepPostgres string

// Save - Salva uma nova ponto, implementacao Postgresql.
func (pontoRepPostgres) Save(tx *sql.Tx, entidade dominio.Ponto) (id int, errDomin *dominio.Erro) {
	return 0, nil
}

// Update - Atualiza um ponto, implementacao Postgresql.
func (pontoRepPostgres) Update(tx *sql.Tx, entidade dominio.Ponto) *dominio.Erro {
	return nil
}

// Delete - Deleta um ponto, implementacao Postgresql.
func (pontoRepPostgres) Delete(tx *sql.Tx, entidade dominio.Ponto) *dominio.Erro {
	return nil
}

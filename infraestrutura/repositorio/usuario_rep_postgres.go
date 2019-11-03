package repositorio

import (
	"database/sql"

	"github.com/gabrielbo1/kronos/dominio"
)

// UsuarioRepPostgres - Tipo Repositorio Usuario para Postgresql.
type usuarioRepPostgres string

// Save - Salva uma nova usuario, implementacao Postgresql.
func (usuarioRepPostgres) Save(tx *sql.Tx, entidade dominio.Usuario) (int, *dominio.Erro) {
	return 0, nil
}

// Update - Atualiza um usuario, implementacao Postgresql.
func (usuarioRepPostgres) Update(tx *sql.Tx, entidade dominio.Usuario) *dominio.Erro {
	return nil
}

// Delete - Deleta uma usuario, implementacao Postgresql.
func (usuarioRepPostgres) Delete(tx *sql.Tx, entidade dominio.Usuario) *dominio.Erro {
	return nil
}

// FindAll - Busca todos usuarios cadastrados, implementacao Postgresql.
func (usuarioRepPostgres) FindAll(tx *sql.Tx) (entidades []dominio.Usuario, erro *dominio.Erro) {
	return nil, nil
}

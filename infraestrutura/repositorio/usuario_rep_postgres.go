package repositorio

import (
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/gabrielbo1/kronos/dominio"
)

// UsuarioRepPostgres - Tipo Repositorio Usuario para Postgresql.
type usuarioRepPostgres string

// Save - Salva uma nova usuario, implementacao Postgresql.
func (usuarioRepPostgres) Save(tx *sql.Tx, entidade dominio.Usuario) (int, *dominio.Erro) {
	sqlInsert := `INSERT INTO usuario(nome, login, senha, acesso) VALUES ($1, $2, $3, $4) RETURNING ID`
	stmt, err := prepararStmt(ctx, tx, "usuarioRepPostgres", "Save", sqlInsert)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var jsonBytes []byte
	var errJson error
	if jsonBytes, errJson = json.Marshal(&entidade.Acesso); err != nil {
		return 0, &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro usuarioRepPostgres função Save", Err: errJson}
	}

	id := 0
	if ok, errDomin := scanParamStmt("usuarioRepPostgres", "Save", stmt, func(stmt *sql.Stmt) error {
		return stmt.QueryRowContext(ctx, &entidade.Nome, &entidade.Login, fmt.Sprintf("%x", sha256.Sum256([]byte(entidade.Senha))), string(jsonBytes)).Scan(&id)
	}); !ok {
		return 0, errDomin
	}
	return id, nil
}

// Update - Atualiza um usuario, implementacao Postgresql.
func (usuarioRepPostgres) Update(tx *sql.Tx, entidade dominio.Usuario) *dominio.Erro {
	sqlUptade := "UPDATE usuario SET nome=$1, login=$2, acesso=$3 WHERE id=$4"
	stmt, err := prepararStmt(ctx, tx, "usuarioRepPostgres", "Update", sqlUptade)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var jsonBytes []byte
	var errJson error
	if jsonBytes, errJson = json.Marshal(&entidade.Acesso); err != nil {
		return &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro usuarioRepPostgres função Save", Err: errJson}
	}

	if ok, errDomain := scanParamStmt("usuarioRepPostgres", "Update", stmt, func(stmt *sql.Stmt) error {
		_, err := stmt.ExecContext(ctx, &entidade.Nome, entidade.Login, string(jsonBytes), entidade.ID)
		return err
	}); !ok {
		return errDomain
	}
	return nil
}

// Delete - Deleta uma usuario, implementacao Postgresql.
func (usuarioRepPostgres) Delete(tx *sql.Tx, entidade dominio.Usuario) *dominio.Erro {
	return delete(tx, "usuarioRepPostgres", "usuario", entidade.ID)
}

// FindAll - Busca todos usuarios cadastrados, implementacao Postgresql.
func (usuarioRepPostgres) FindAll(tx *sql.Tx) (entidades []dominio.Usuario, erro *dominio.Erro) {
	rows, errTx := tx.QueryContext(ctx, "SELECT id, nome, login, senha, acesso FROM usuario")
	if errTx != nil {
		return []dominio.Usuario{}, &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro FindAll usuarioRepPostgres função FindAll", Err: errTx}
	}
	usuarios, err := parseUsuarioPostgres(rows)
	if err != nil {
		return []dominio.Usuario{}, &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro FindAll usuarioRepPostgres função FindAll", Err: errTx}
	}
	return usuarios, nil
}

func (usuarioRepPostgres) Login(tx *sql.Tx, login, senha string) (dominio.Usuario, *dominio.Erro) {
	sqlQuery := "SELECT id, nome, login, senha, acesso FROM usuario WHERE login = $1 AND senha = $2"
	stmt, errDomin := prepararStmt(ctx, tx, "Login", "FindById", sqlQuery)
	if errDomin != nil {
		return dominio.Usuario{}, errDomin
	}
	defer stmt.Close()

	rows, errTx := stmt.QueryContext(ctx, login, fmt.Sprintf("%x", sha256.Sum256([]byte(senha))))
	if errTx != nil {
		return dominio.Usuario{}, &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro Login usuarioRepPostgres função Login", Err: errTx}
	}

	usuarios, err := parseUsuarioPostgres(rows)
	if err != nil {
		return dominio.Usuario{}, &dominio.Erro{Codigo: "SQLUTIL_REP20", Mensagem: "Erro Login usuarioRepPostgres função Login", Err: err}
	}
	if len(usuarios) != 0 {
		return usuarios[0], nil
	}
	return dominio.Usuario{}, nil
}

func parseUsuarioPostgres(rows *sql.Rows) ([]dominio.Usuario, error) {
	defer rows.Close()
	var result []dominio.Usuario
	for rows.Next() {
		usu := dominio.Usuario{}
		jsonAcesso := ""
		if err := rows.Scan(&usu.ID, &usu.Nome, &usu.Login, &usu.Senha, &jsonAcesso); err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(jsonAcesso), &usu.Acesso); err != nil {
			return nil, err
		}
		result = append(result, usu)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return result, nil
}

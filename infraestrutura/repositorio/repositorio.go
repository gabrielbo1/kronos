package repositorio

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/gabrielbo1/kronos/infraestrutura"

	"github.com/gabrielbo1/kronos/dominio"
)

// SGDB - Define o tipo de banco de dados.
type SGDB string

//POSTGRES - Banco PostgreSQL.
const POSTGRES SGDB = "POSTGRES"

//DB - Conexao com banco de dados.
var DB *sql.DB

func stringConexaoPostgres(confg infraestrutura.Configuracao) string {
	connString := "host="
	connString += confg.IPBanco
	connString += " user="
	connString += confg.UsuarioBanco
	connString += " dbname="
	connString += confg.NomeBanco
	connString += "  password='"
	connString += confg.SenhaBanco + "'"
	connString += " sslmode=disable"
	return connString
}

// BuscaConexao - Busca conexacao com banco de dados.
func buscaConexao() (DB *sql.DB, errDomin *dominio.Erro) {
	if DB == nil {
		switch SGDB(infraestrutura.Config.Banco) {
		case POSTGRES:
			var err error
			DB, err = sql.Open("postgres", stringConexaoPostgres(infraestrutura.Config))
			if err != nil {
				log.Fatal(err)
				return nil, &dominio.Erro{"CON_10", "Erro banco FindDb", err}
			}
			return DB, errDomin
		}
	}
	return nil, &dominio.Erro{Codigo: "REPOSITORIO10", Mensagem: "Banco de dados nao suportado ou erro na conexao."}
}

// ShcemaUpdate -  Executa migracao.
func ShcemaUpdate(diretorioScripts string) *dominio.Erro {
	var e *dominio.Erro
	switch SGDB(infraestrutura.Config.Banco) {
	case POSTGRES:
		if DB, e = buscaConexao(); e != nil {
			return e
		}

		files, err := ioutil.ReadDir(diretorioScripts)
		if err != nil {
			log.Fatal(err)
		}

		var schemas []Schema
		for _, f := range files {
			body, err := ioutil.ReadFile(filepath.Join(diretorioScripts, f.Name()))
			if err != nil {
				log.Fatalf("Não foi possivel ler arquivo de scripts_sql: %v", err)
				return &dominio.Erro{Codigo: "REPOSITORIO20",
					Mensagem: fmt.Sprintf("Não foi possivel ler arquivo de scripts_sql: %v", f.Name())}
			}
			schemas = append(schemas, Schema{0, f.Name(), string(body), false})
		}
		InitSchemaBd(schemas)
		return nil
	}
	return nil
}

// Transact - Permite o controle e execução de várias transações
// de maneira que as transações fiquem aninhadas umas
// as outras e qualquer problema a função identifica
// e realiza o roolback da transação.
// func (s Service) DoSomething() error {
//    return Transact(s.db, func (tx *scripts_sql.Tx) error {
//        if _, err := tx.Exec(...); err != nil {
//            return err
//        }
//        if _, err := tx.Exec(...); err != nil {
//            return err
//        }
//    })
// }
func Transact(db *sql.DB, txFunc func(*sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return err
}

//Pagina - Define estrutura para consultas paginadas.
type Pagina struct {
	NumPagina     int         `json:"numPagina" example:"1"`
	QtdPorPagina  int         `json:"qtdPagina" example:"10"`
	TotalRegistro int         `json:"totalRegistro" example:"1000"`
	TotalPagina   int         `json:"totalPagina" example:"100"`
	Conteudo      interface{} `json:"conteudo"`
}

// EmpresaRepositorio - Define operacoes a serem realizadas
// com a entidade empresa.
type EmpresaRepositorio interface {
	Save(tx *sql.Tx, entidade dominio.Empresa) *dominio.Erro

	Update(tx *sql.Tx, entidade dominio.Empresa) *dominio.Erro

	Delete(tx *sql.Tx, entidade dominio.Empresa) *dominio.Erro

	FindAll(tx *sql.Tx) (entidades []dominio.Empresa, erro *dominio.Erro)
}

// NewEmpresaRepositorio - Retorna repositorio de empresa.
func NewEmpresaRepositorio() EmpresaRepositorio {
	switch SGDB(infraestrutura.Config.Banco) {
	case POSTGRES:
		var rep empresaRepPostgres
		rep = "empresa_rep_postgres"
		return rep
	}
	return nil
}

// UsuarioRepositorio - Define operacoes a serem realizadas
// com a entidade Usuario.
type UsuarioRepositorio interface {
	Save(tx *sql.Tx, entidade dominio.Usuario) (int, *dominio.Erro)

	Update(tx *sql.Tx, entidade dominio.Usuario) *dominio.Erro

	Delete(tx *sql.Tx, entidade dominio.Usuario) *dominio.Erro

	FindAll(tx *sql.Tx) (entidades []dominio.Usuario, erro *dominio.Erro)

	Login(tx *sql.Tx, login, senha string) (dominio.Usuario, *dominio.Erro)
}

//NewUsuarioRepositorio - Retorna repositorio de usuario.
func NewUsuarioRepositorio() UsuarioRepositorio {
	switch SGDB(infraestrutura.Config.Banco) {
	case POSTGRES:
		var rep usuarioRepPostgres
		rep = "usuario_rep_postgres"
		return rep
	}
	return nil
}

// RotinaRepositorio - Define operacoes a serem realizadas
// com a entidade Rotina
type RotinaRepositorio interface {
	Save(tx *sql.Tx, entidade dominio.Rotina) (int, *dominio.Erro)

	Update(tx *sql.Tx, entidade dominio.Rotina) *dominio.Erro

	Delete(tx *sql.Tx, entidade dominio.Rotina) *dominio.Erro

	FindAll(tx *sql.Tx) (entidades []dominio.Rotina, erro *dominio.Erro)
}

// NewRotinaRepositorio - Retorna repositorio de rotina.
func NewRotinaRepositorio() RotinaRepositorio {
	switch SGDB(infraestrutura.Config.Banco) {
	case POSTGRES:
		var rep rotinaRepPostgres
		rep = "rotina_rep_postgres"
		return rep
	}
	return nil
}

// PontoRepositorio - Define operacoes a serem realizadas
// com a entidade Ponto.
type PontoRepositorio interface {
	Save(tx *sql.Tx, entidade dominio.Ponto) (int, *dominio.Erro)

	Update(tx *sql.Tx, entidade dominio.Ponto) *dominio.Erro

	Delete(tx *sql.Tx, entidade dominio.Ponto) *dominio.Erro
}

// NewPontoRepositorio - Retorna repositorio de ponto.
func NewPontoRepositorio() PontoRepositorio {
	switch SGDB(infraestrutura.Config.Banco) {
	case POSTGRES:
		var rep pontoRepPostgres
		rep = "ponto_rep_postgres"
		return rep
	}
	return nil
}

// AtendimentoRepositorio - Define operacoes a serem realizadas
// com a entidade Ponto.
type AtendimentoRepositorio interface {
	Save(tx *sql.Tx, entidade dominio.Atendimento) (int, *dominio.Erro)

	Update(tx *sql.Tx, entidade dominio.Atendimento) *dominio.Erro

	Delete(tx *sql.Tx, entidade dominio.Atendimento) *dominio.Erro
}

// NewRepositorioAtendimento -  Retorna repositorio de atendimento
func NewRepositorioAtendimento() AtendimentoRepositorio {
	switch SGDB(infraestrutura.Config.Banco) {
	case POSTGRES:
		var rep atendimentoRepPostgres
		rep = "atendimento_rep_postgres"
		return rep
	}
	return nil
}

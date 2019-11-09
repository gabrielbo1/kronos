package repositorio

import (
	"context"
	"database/sql"
	"log"
	"sort"
	"strings"

	"github.com/gabrielbo1/kronos/infraestrutura"

	"github.com/gabrielbo1/kronos/dominio"
)

var ctx = context.Background()

// Schema - Define estrutura para controle de versao do banco.
type Schema struct {
	ID        int
	Nome      string
	SQL       string
	Executado bool
}

func parseASCIINumber(name string) int {
	i := 0
	for k := range strings.Split(name, "") {
		i += int(k)
	}
	return i
}

// SortSchema - Ordenacao de acordo com os nomes e numeros.
// atribuidos aos arquivos de controle de versao.
func sortSchema(schema1, schema2 Schema) bool {
	return parseASCIINumber(schema1.Nome) < parseASCIINumber(schema1.Nome)
}

// CreateIfNotExistis -
func CreateIfNotExistis(tx *sql.Tx) *dominio.Erro {
	sql := "CREATE SEQUENCE  IF NOT EXISTS schema_controle_id_seq "
	sql += "INCREMENT 1  MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1;"
	sql += "CREATE TABLE  IF NOT EXISTS SCHEMA_CONTROLE ( "
	sql += "ID BIGINT DEFAULT nextval('schema_controle_id_seq') PRIMARY KEY, "
	sql += "NOME VARCHAR (20) NOT NULL, "
	sql += "EXECUTADO BOOLEAN NOT NULL DEFAULT FALSE "
	sql += "); "

	_, err := tx.Exec(sql)
	if err != nil {
		log.Println(err)
		return &dominio.Erro{"SHCMEA_REP05", "Erro banco CREATE TABLE SCHEMA", err}
	}
	return nil
}

func save(tx *sql.Tx, schema Schema) *dominio.Erro {
	query := inserIntoSchema(schema)
	_, err := tx.Exec(query)
	if err != nil {
		log.Println(err)
		return &dominio.Erro{"SHCMEA_REP10", "Erro banco Save", err}
	}
	return nil
}

func execSchmea(tx *sql.Tx, schema Schema) *dominio.Erro {
	sql := strings.Replace(schema.SQL, "\n", " ", -1)
	_, err := tx.Exec(sql)
	if err != nil {
		log.Println("Erro ao executar schema, " + schema.Nome)
		return &dominio.Erro{"SHCMEA_REP20", "Erro ao executar schema, " + schema.Nome, err}
	}
	log.Println("Sucesso ao executar schema, " + schema.Nome)
	save(tx, Schema{0, schema.Nome, "", true})
	return nil
}

// FindAll -
func FindAll(tx *sql.Tx) (entidades []Schema, erro *dominio.Erro) {
	var schemas []Schema
	sql := "SELECT ID, NOME, EXECUTADO FROM SCHEMA_CONTROLE"
	stmt, err := tx.PrepareContext(ctx, sql)
	if err != nil {
		log.Println(err)
		return nil, &dominio.Erro{"SHCMEA_REP30", "Erro ao buscar schemas, ", err}
	}
	if rows, err := stmt.QueryContext(ctx); err == nil {
		defer rows.Close()
		for rows.Next() {
			var schema Schema = Schema{}
			if err = rows.Scan(&schema.ID, &schema.Nome, &schema.Executado); err == nil {
				schemas = append(schemas, schema)
			}
		}
	} else if err != nil {
		return nil, &dominio.Erro{"SHCMEA_REP30", "Erro ao buscar schemas, ", err}
	}
	return schemas, nil
}

func inserIntoSchema(schema Schema) string {
	return "INSERT INTO SCHEMA_CONTROLE (NOME, EXECUTADO) VALUES(" + StringToString(schema.Nome) + "," + BooleanToString(schema.Executado) + ")"
}

func createIfNotExistis(tx *sql.Tx) *dominio.Erro {
	sql := "CREATE SEQUENCE  IF NOT EXISTS schema_controle_id_seq "
	sql += "INCREMENT 1  MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1;"
	sql += "CREATE TABLE  IF NOT EXISTS SCHEMA_CONTROLE ( "
	sql += "ID BIGINT DEFAULT nextval('schema_controle_id_seq') PRIMARY KEY, "
	sql += "NOME VARCHAR (20) NOT NULL, "
	sql += "EXECUTADO BOOLEAN NOT NULL DEFAULT FALSE "
	sql += "); "

	_, err := tx.Exec(sql)
	if err != nil {
		log.Println(err)
		return &dominio.Erro{"SHCMEA_REP05", "Erro banco CREATE TABLE SCHEMA", err}
	}
	return nil
}

// InitSchemaBd - Realiza operacao de migracao da base dados.
func InitSchemaBd(schemas []Schema) {
	var schemasBd []Schema
	var err *dominio.Erro
	var db *sql.DB

	log.Println("INICIANDO MIGRACAO BASE DE DADOS: " + infraestrutura.Config.IPBanco + " - " + infraestrutura.Config.NomeBanco)

	if db = DB; db != nil {
		if errTx := Transact(db, func(tx *sql.Tx) error {
			err = createIfNotExistis(tx)
			return nil
		}); errTx != nil {
			log.Fatal(errTx)
		}
	}

	if db = DB; db != nil {
		if errTx := Transact(db, func(tx *sql.Tx) error {
			schemasBd, err = FindAll(tx)
			return nil
		}); errTx != nil {
			log.Fatal(errTx)
		}
	}

	if schemasBd != nil && len(schemasBd) <= len(schemas) {
		sort.SliceStable(schemasBd, func(i, j int) bool {
			return sortSchema(schemasBd[i], schemasBd[j])
		})
		sort.SliceStable(schemas, func(i, j int) bool {
			return sortSchema(schemas[i], schemas[j])
		})
		i := 0
		for ; i < len(schemasBd) && i < len(schemas); i++ {
			if schemasBd[i].Nome == schemas[i].Nome && !schemasBd[i].Executado {
				if errTx := Transact(db, func(tx *sql.Tx) error {
					err = execSchmea(tx, schemas[i])
					return nil
				}); errTx != nil {
					log.Fatal(errTx)
				}
			}
		}

		for ; i < len(schemas); i++ {
			if errTx := Transact(db, func(tx *sql.Tx) error {
				err = execSchmea(tx, schemas[i])
				return nil
			}); errTx != nil {
				log.Fatal(errTx)
			}
		}
	} else {
		for i := 0; i < len(schemas); i++ {
			if errTx := Transact(db, func(tx *sql.Tx) error {
				err = execSchmea(tx, schemas[i])
				return nil
			}); errTx != nil {
				log.Fatal(errTx)
			}
		}
	}
	if err != nil {
		if err.Err != nil {
			log.Fatalf("%v, %v", err.Codigo, err.Mensagem)
		} else {
			log.Fatalf("%v, %v, %v", err.Codigo, err.Mensagem, err.Err)
		}
	}
}

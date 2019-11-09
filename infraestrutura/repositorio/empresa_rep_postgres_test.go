package repositorio

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/gabrielbo1/kronos/dominio"

	txdb "github.com/DATA-DOG/go-txdb"
	"github.com/gabrielbo1/kronos/infraestrutura"
	_ "github.com/lib/pq"
)

func init() {
	infraestrutura.Config.NomeBanco = "kronostest"
	txdb.Register("psql_txdb", "postgres", stringConexaoPostgres(infraestrutura.Config))
}

func preparePostgresDB(t *testing.T) (db *sql.DB, teardown func() error) {
	dbName := fmt.Sprintf("db_%d", time.Now().UnixNano())
	db, err := sql.Open("psql_txdb", dbName)

	if err != nil {
		log.Fatalf("open postgresql connection: %s", err)
	}
	DB = db
	if e := ShcemaUpdate("./script_postgressql"); e != nil {
		t.Fatal(e)
	}
	return db, db.Close
}

func TestEmpresaRepPostgres_Save(t *testing.T) {
	db, destruirBd := preparePostgresDB(t)
	defer destruirBd()
	defer db.Close()

	rep := NewEmpresaRepositorio()
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	emp := dominio.EmpresaMock()
	if errDomin := rep.Save(tx, emp); errDomin != nil {
		t.Error(errDomin)
		t.Fail()
	}
	t.Log(emp)
}

func TestEmpresaRepPostgres_Update(t *testing.T) {
	db, destruirBd := preparePostgresDB(t)
	defer destruirBd()
	defer db.Close()

	rep := NewEmpresaRepositorio()
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	emp := dominio.EmpresaMock()

	errDomin := rep.Save(tx, emp)
	if errDomin != nil {
		t.Error(errDomin)
		t.Fail()
	}

	emp.Nome = "abc"
	if errDomin = rep.Update(tx, emp); errDomin != nil {
		t.Error(errDomin)
		t.Fail()
	}

}

func TestEmpresaRepPostgres_Delete(t *testing.T) {
	db, destruirBd := preparePostgresDB(t)
	defer destruirBd()
	defer db.Close()

	rep := NewEmpresaRepositorio()
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	emp := dominio.EmpresaMock()

	errDomin := rep.Save(tx, emp)
	if errDomin != nil {
		t.Error(errDomin)
		t.Fail()
	}

	if errDomin = rep.Delete(tx, emp); errDomin != nil {
		t.Error(errDomin)
		t.Fail()
	}
}

func TestEmpresaRepPostgres_FindAll(t *testing.T) {
	db, destruirBd := preparePostgresDB(t)
	defer destruirBd()
	defer db.Close()

	rep := NewEmpresaRepositorio()
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		if errDomin := rep.Save(tx, dominio.EmpresaMock()); errDomin != nil {
			t.Error(errDomin)
			t.Fail()
		}
	}

	empresas, errDomin := rep.FindAll(tx)
	if errDomin != nil || len(empresas) != 10 {
		t.Error(errDomin)
		t.Fail()
	}

	t.Log(empresas)
}

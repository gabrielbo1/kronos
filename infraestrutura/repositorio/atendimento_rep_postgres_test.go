package repositorio

import (
	"database/sql"
	"github.com/gabrielbo1/kronos/dominio"
	"testing"
)

func criarEmpresa(t *testing.T, tx *sql.Tx) dominio.Empresa {
	repEmpresa := NewEmpresaRepositorio()

	emp := dominio.EmpresaMock()
	var errDominio *dominio.Erro

	emp.ID, errDominio = repEmpresa.Save(tx, emp)
	if errDominio != nil {
		t.Error(errDominio)
		t.Fail()
	}
	return emp
}

func TestAtendimentoRepPostgres_Save(t *testing.T) {
	db, destruirBd := preparePostgresDB(t)
	defer destruirBd()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	usu := criaUsuario(t, tx)
	emp := criarEmpresa(t, tx)

	aten := dominio.AtendimentoMock()
	aten.Usuario = usu
	aten.Cliente = emp

	rep := NewRepositorioAtendimento()

	if _, errDomin := rep.Save(tx, aten); errDomin != nil {
		t.Error(errDomin)
		t.Fail()
	}
}

func TestAtendimentoRepPostgres_Update(t *testing.T) {
	db, destruirBd := preparePostgresDB(t)
	defer destruirBd()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	usu := criaUsuario(t, tx)
	emp := criarEmpresa(t, tx)

	aten := dominio.AtendimentoMock()
	aten.Usuario = usu
	aten.Cliente = emp

	rep := NewRepositorioAtendimento()
	var errDominio *dominio.Erro

	if aten.ID, errDominio = rep.Save(tx, aten); errDominio != nil {
		t.Error(errDominio)
		t.Fail()
	}

	aten.Observacao = "abcd"
	if errDominio = rep.Update(tx, aten); errDominio != nil {
		t.Error(errDominio)
		t.Fail()
	}
}

func TestAtendimentoRepPostgres_Delete(t *testing.T) {
	db, destruirBd := preparePostgresDB(t)
	defer destruirBd()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	usu := criaUsuario(t, tx)
	emp := criarEmpresa(t, tx)

	aten := dominio.AtendimentoMock()
	aten.Usuario = usu
	aten.Cliente = emp

	rep := NewRepositorioAtendimento()
	var errDominio *dominio.Erro

	if aten.ID, errDominio = rep.Save(tx, aten); errDominio != nil {
		t.Error(errDominio)
		t.Fail()
	}

	if errDominio = rep.Delete(tx, aten); errDominio != nil {
		t.Error(errDominio)
		t.Fail()
	}
}

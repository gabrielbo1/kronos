package repositorio

import (
	"database/sql"
	"github.com/gabrielbo1/kronos/dominio"
	"testing"
	"time"
)

func criaUsuario(t *testing.T, tx *sql.Tx) dominio.Usuario {
	repAcesso := NewRotinaRepositorio()
	rot := dominio.RotinaMock()
	var errDomin *dominio.Erro
	rot.ID, errDomin = repAcesso.Save(tx, rot)

	if errDomin != nil {
		t.Error(errDomin)
		t.Fail()
	}

	usu := dominio.UsuarioMock()
	usu.Acesso = []dominio.Acesso{}
	usu.Acesso = append(usu.Acesso, dominio.AcessoMock())
	usu.Acesso[0].Rotina = rot

	rep := NewUsuarioRepositorio()
	if usu.ID, errDomin = rep.Save(tx, usu); errDomin != nil {
		t.Error(errDomin)
		t.Fail()
	}
	return usu
}

func TestPontoRepPostgres_Save(t *testing.T) {
	db, destruirBd := preparePostgresDB(t)
	defer destruirBd()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	usu := criaUsuario(t, tx)

	rep := NewPontoRepositorio()
	ponto := dominio.PontoMock()
	ponto.Usuario = usu

	if _, errDomin := rep.Save(tx, ponto); errDomin != nil {
		t.Error(errDomin)
		t.Fail()
	}
}

func TestPontoRepPostgres_Update(t *testing.T) {
	db, destruirBd := preparePostgresDB(t)
	defer destruirBd()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	usu := criaUsuario(t, tx)

	rep := NewPontoRepositorio()
	ponto := dominio.PontoMock()
	ponto.Usuario = usu

	var errDomin *dominio.Erro
	if ponto.ID, errDomin = rep.Save(tx, ponto); errDomin != nil {
		t.Error(errDomin)
		t.Fail()
	}

	ponto.Data = time.Now().Add(time.Hour).Format(time.RFC822)
	if errDomin = rep.Update(tx, ponto); errDomin != nil {
		t.Error(errDomin)
		t.Fail()
	}
}

func TestPontoRepPostgres_Delete(t *testing.T) {
	db, destruirBd := preparePostgresDB(t)
	defer destruirBd()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	usu := criaUsuario(t, tx)

	rep := NewPontoRepositorio()
	ponto := dominio.PontoMock()
	ponto.Usuario = usu

	var errDomin *dominio.Erro
	if ponto.ID, errDomin = rep.Save(tx, ponto); errDomin != nil {
		t.Error(errDomin)
		t.Fail()
	}

	if errDomin = rep.Delete(tx, ponto); errDomin != nil {
		t.Error(errDomin)
		t.Fail()
	}
}

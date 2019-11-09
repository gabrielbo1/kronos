package repositorio

import (
	"testing"

	"github.com/gabrielbo1/kronos/dominio"
)

func TestSaveUsuario(t *testing.T) {
	db, destruirBd := preparePostgresDB(t)
	defer destruirBd()
	defer db.Close()

	repAcesso := NewRotinaRepositorio()
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

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
}

func TestUpdateUsuario(t *testing.T) {
	db, destruirBd := preparePostgresDB(t)
	defer destruirBd()
	defer db.Close()

	repAcesso := NewRotinaRepositorio()
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

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

	usu.Login = "abc"
	if errDomin := rep.Update(tx, usu); errDomin != nil {
		t.Error(errDomin)
		t.Fail()
	}
}

func TestDeleteUsuario(t *testing.T) {
	db, destruirBd := preparePostgresDB(t)
	defer destruirBd()
	defer db.Close()

	repAcesso := NewRotinaRepositorio()
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

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

	if errDomin := rep.Delete(tx, usu); errDomin != nil {
		t.Error(errDomin)
		t.Fail()
	}
}

func TestFindAllUsuario(t *testing.T) {
	db, destruirBd := preparePostgresDB(t)
	defer destruirBd()
	defer db.Close()

	repAcesso := NewRotinaRepositorio()
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

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
	for i := 0; i < 10; i++ {
		if usu.ID, errDomin = rep.Save(tx, usu); errDomin != nil {
			t.Error(errDomin)
			t.Fail()
		}
	}

	usuarios, errDomin := rep.FindAll(tx)
	if errDomin != nil || len(usuarios) != 10 {
		t.Error(errDomin)
		t.Fail()
	}
	t.Log(usuarios)
}

func TestUsuarioRepPostgres_Login(t *testing.T) {
	db, destruirBd := preparePostgresDB(t)
	defer destruirBd()
	defer db.Close()

	repAcesso := NewRotinaRepositorio()
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

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

	if usu, err := rep.Login(tx, usu.Login, usu.Senha); err != nil || usu.ID == 0 {
		t.Error(errDomin)
		t.Fail()
	}

	t.Log(usu)
}

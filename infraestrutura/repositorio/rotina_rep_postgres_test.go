package repositorio

import (
	"testing"

	"github.com/gabrielbo1/kronos/dominio"
)

func TestRotinaRepPostgres_Save(t *testing.T) {
	db, destruirBd := preparePostgresDB(t)
	defer destruirBd()
	defer db.Close()

	rep := NewRotinaRepositorio()
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	if _, errDomin := rep.Save(tx, dominio.RotinaMock()); errDomin != nil {
		t.Error(errDomin)
		t.Fail()
	}
}

func TestRotinaRepPostgres_Update(t *testing.T) {
	db, destruirBd := preparePostgresDB(t)
	defer destruirBd()
	defer db.Close()

	rep := NewRotinaRepositorio()
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	rot := dominio.RotinaMock()
	var errDomin *dominio.Erro
	rot.ID, errDomin = rep.Save(tx, rot)

	if errDomin != nil {
		t.Error(errDomin)
		t.Fail()
	}

	rot.Rotina = "abc"
	if errDomin = rep.Update(tx, rot); errDomin != nil {
		t.Error(errDomin)
		t.Fail()
	}
}

func TestRotinaRepPostgres_Delete(t *testing.T) {
	db, destruirBd := preparePostgresDB(t)
	defer destruirBd()
	defer db.Close()

	rep := NewRotinaRepositorio()
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	rot := dominio.RotinaMock()
	var errDomin *dominio.Erro
	rot.ID, errDomin = rep.Save(tx, rot)

	if errDomin != nil {
		t.Error(errDomin)
		t.Fail()
	}

	if errDomin = rep.Delete(tx, rot); errDomin != nil {
		t.Error(errDomin)
		t.Fail()
	}
}

func TestRotinaRepPostgres_FindAll(t *testing.T) {
	db, destruirBd := preparePostgresDB(t)
	defer destruirBd()
	defer db.Close()

	rep := NewRotinaRepositorio()
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		if _, errDomin := rep.Save(tx, dominio.RotinaMock()); errDomin != nil {
			t.Error(errDomin)
			t.Fail()
		}
	}

	rotinas, errDomin := rep.FindAll(tx)
	if errDomin != nil || len(rotinas) != 10 {
		t.Error(errDomin)
		t.Fail()
	}

	t.Log(rotinas)
}

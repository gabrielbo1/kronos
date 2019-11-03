package dominio

import (
	"testing"
	"time"
)

func TestNewPonto(t *testing.T) {
	p := Ponto{}
	var e *Erro

	if e = NewPonto(&p); e == nil || e.Codigo != "PONTO10" {
		t.Log("TestNewPonto - PONTO10")
		t.Fail()
	}

	p.Usuario = UsuarioMock()
	if e = NewPonto(&p); e == nil || e.Codigo != "PONTO20" {
		t.Log("TestNewPonto - PONTO20")
		t.Fail()
	}

	ontem := time.Now().Add(time.Hour * 24 * -2)
	p.Data = ontem.Format(time.RFC3339)
	if e = NewPonto(&p); e == nil || e.Codigo != "PONTO30" {
		t.Log("TestNewPonto - PONTO30")
		t.Fail()
	}

	amanha := time.Now().Add(time.Hour * 2 * 24)
	p.Data = amanha.Format(time.RFC3339)
	if e = NewPonto(&p); e == nil || e.Codigo != "PONTO30" {
		t.Log("TestNewPonto - PONTO30")
		t.Fail()
	}

	p.Data = time.Now().Format(time.RFC3339)
	if e = NewPonto(&p); e != nil {
		t.Log("TestNewPonto")
		t.Fail()
	}
}

func TestNewPontoAvulso(t *testing.T) {
	p := Ponto{}
	var e *Erro

	if e = NewPontoAvulso(&p); e == nil || e.Codigo != "PONTO10" {
		t.Log("TestNewPonto - PONTO10")
		t.Fail()
	}

	p.Usuario = UsuarioMock()
	if e = NewPontoAvulso(&p); e == nil || e.Codigo != "PONTO20" {
		t.Log("TestNewPonto - PONTO20")
		t.Fail()
	}

	ontem := time.Now().Add(time.Hour * 24 * -2)
	amanha := time.Now().Add(time.Hour * 2 * 24)

	p.Data = ontem.Format(time.RFC3339)
	if e = NewPontoAvulso(&p); e != nil {
		t.Log("TestNewPonto")
		t.Fail()
	}

	p.Data = amanha.Format(time.RFC3339)
	if e = NewPontoAvulso(&p); e != nil {
		t.Log("TestNewPonto")
		t.Fail()
	}
}

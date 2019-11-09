package dominio

import "testing"

func TestNewRotina(t *testing.T) {
	r := Rotina{}
	var e *Erro

	if e = NewRotina(&r); e == nil || e.Codigo != "ROTINA10" {
		t.Log("TestNewRotina - ROTINA10")
		t.Fail()
	}

	r.Rotina = "/CadUsuario"
	if e = NewRotina(&r); e != nil {
		t.Log("TestNewRotina")
		t.Fail()
	}
}

func TestNewAcesso(t *testing.T) {
	a := Acesso{}
	var e *Erro

	a.Rotina = RotinaMock()

	if e = NewAcesso(&a); e == nil || e.Codigo != "ACESSO10" {
		t.Log("TestNewAcesso - ACESSO10")
		t.Fail()
	}

	a.Criar = true
	if e = NewAcesso(&a); e != nil {
		t.Log("TestNewAcesso")
		t.Fail()
	}
}

func TestNewUsuario(t *testing.T) {
	usu := Usuario{}
	var e *Erro

	if e = NewUsuario(&usu); e == nil || e.Codigo != "USUARIO10" {
		t.Log("TestNewUsuario - USUARIO10")
		t.Fail()
	}

	usu.Nome = "gabriel"
	if e = NewUsuario(&usu); e == nil || e.Codigo != "USUARIO20" {
		t.Log("TestNewUsuario - USUARIO20")
		t.Fail()
	}

	usu.Nome = "gabriel"
	usu.Login = "gbo"
	if e = NewUsuario(&usu); e == nil || e.Codigo != "USUARIO30" {
		t.Log("TestNewUsuario - USUARIO30")
		t.Fail()
	}

	usu.Nome = "gabriel"
	usu.Login = "gbo"
	usu.Senha = "pass123"
	if e = NewUsuario(&usu); e == nil || e.Codigo != "USUARIO40" {
		t.Log("TestNewUsuario - USUARIO40")
		t.Fail()
	}

	usu.Acesso = append(usu.Acesso, AcessoMock())
	if e = NewUsuario(&usu); e != nil {
		t.Log("TestNewUsuario")
		t.Fail()
	}
}

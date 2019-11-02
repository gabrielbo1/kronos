package dominio

import "testing"

func TestNewEmpresa(t *testing.T) {
	emp := Empresa{}
	var e *Erro

	if e = NewEmpresa(&emp); e == nil || e.Codigo != "EMPRESA10" {
		t.Fail()
	}

	emp.Nome = "Nome Compania"
	if e = NewEmpresa(&emp); e != nil {
		t.Logf("%v - %v\n", e.Codigo, e.Mensagem)
		t.Fail()
	}
}

package dominio

import (
	"testing"
	"time"
)

func TestNewAtendimento(t *testing.T) {
	aten := Atendimento{}
	var e *Erro

	if e = NewAtendimento(&aten); e == nil || e.Codigo != "ATENDIMENTO10" {
		t.Log("TestNewAtendimento - ATENDIMENTO10")
		t.Fail()
	}

	aten.Usuario = UsuarioMock()
	if e = NewAtendimento(&aten); e == nil || e.Codigo != "ATENDIMENTO20" {
		t.Log("TestNewAtendimento - ATENDIMENTO20")
		t.Fail()
	}

	aten.Cliente = EmpresaMock()
	if e = NewAtendimento(&aten); e == nil || e.Codigo != "ATENDIMENTO30" {
		t.Log("TestNewAtendimento - ATENDIMENTO30")
		t.Fail()
	}

	aten.HorariosAtendimento = append(aten.HorariosAtendimento, Intervalo{DataInicio: ""})
	if e = NewAtendimento(&aten); e == nil || e.Codigo != "ATENDIMENTO40" {
		t.Log("TestNewAtendimento - ATENDIMENTO40")
		t.Fail()
	}

	aten.HorariosAtendimento[0].DataInicio = time.Now().Format(time.RFC3339)
	aten.HorariosAtendimento[0].DataFim = "zzz"
	if e = NewAtendimento(&aten); e == nil || e.Codigo != "ATENDIMENTO50" {
		t.Log("TestNewAtendimento - ATENDIMENTO50")
		t.Fail()
	}

	aten.HorariosAtendimento[0].DataInicio = time.Now().Format(time.RFC3339)
	aten.HorariosAtendimento[0].DataFim = ""
	aten.StatusAtendimento = Fechado
	if e = NewAtendimento(&aten); e == nil || e.Codigo != "ATENDIMENTO60" {
		t.Log("TestNewAtendimento - ATENDIMENTO60")
		t.Fail()
	}

	aten.HorariosAtendimento[0].DataInicio = time.Now().Format(time.RFC3339)
	aten.HorariosAtendimento[0].DataFim = ""
	aten.StatusAtendimento = Aberto
	if e = NewAtendimento(&aten); e != nil {
		t.Log("TestNewAtendimento")
		t.Fail()
	}

	aten.StatusAtendimento = Espera
	if e = NewAtendimento(&aten); e != nil {
		t.Log("TestNewAtendimento")
		t.Fail()
	}

	aten.HorariosAtendimento[0].DataFim = time.Now().Add(time.Hour).Format(time.RFC3339)
	aten.StatusAtendimento = Fechado
	if e = NewAtendimento(&aten); e != nil {
		t.Log("TestNewAtendimento")
		t.Fail()
	}
}

func AtendimentoMock() Atendimento {
	aten := Atendimento{ID: 1, Usuario: UsuarioMock(), Cliente: EmpresaMock()}
	aten.HorariosAtendimento = append(aten.HorariosAtendimento, Intervalo{})
	aten.HorariosAtendimento[0].DataInicio = time.Now().Format(time.RFC3339)
	aten.HorariosAtendimento[0].DataFim = time.Now().Add(time.Hour).Format(time.RFC3339)
	aten.StatusAtendimento = Fechado
	return aten
}

package dominio

import "time"

// RotinaMock - Mock rotina.
func RotinaMock() Rotina {
	return Rotina{ID: 1, Rotina: "/CadUsuario"}
}

// AcessoMock - Mock acesso.
func AcessoMock() Acesso {
	return Acesso{Rotina: RotinaMock(), Criar: true}
}

// UsuarioMock - Mock usuario.
func UsuarioMock() Usuario {
	usu := Usuario{
		ID:    1,
		Nome:  "gabriel",
		Login: "gbo",
		Senha: "pass123",
	}
	usu.Acesso = append(usu.Acesso, AcessoMock())
	return usu
}

// EmpresaMock - Mock empresa.
func EmpresaMock() Empresa {
	return Empresa{ID: 1, Nome: "Empresa X", Ativa: true}
}

// PontoMock - Mock ponto.
func PontoMock() Ponto {
	return Ponto{ID: 1, Usuario: UsuarioMock(), Data: time.Now().Format(time.RFC822)}
}

// AtendimentoMock - Mock atendimento.
func AtendimentoMock() Atendimento {
	aten := Atendimento{ID: 1, Usuario: UsuarioMock(), Cliente: EmpresaMock()}
	aten.HorariosAtendimento = append(aten.HorariosAtendimento, Intervalo{})
	aten.HorariosAtendimento[0].DataInicio = time.Now().Format(time.RFC3339)
	aten.HorariosAtendimento[0].DataFim = time.Now().Add(time.Hour).Format(time.RFC3339)
	aten.StatusAtendimento = Fechado
	return aten
}

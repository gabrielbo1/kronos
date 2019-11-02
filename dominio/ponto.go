package dominio

import (
	"time"

	"github.com/jinzhu/now"
)

// Ponto - Entidade para controle de entrada e saida
// dos colaboradores.
type Ponto struct {
	ID      int     `json:"id"`
	Usuario Usuario `json:"usuario"`
	Data    string  `json:"data"`
}

func ontem(data time.Time) time.Time {
	return data.Add(time.Duration(data.Hour()) * time.Hour * -1).Add(time.Minute * -1)
}

func amanha(data time.Time) time.Time {
	return data.Add(time.Duration(24-data.Hour()) * time.Hour).Add(time.Minute)
}

// NewPonto - Cria um novo ponto para registro de horarios de entrada e saida.
func NewPonto(ponto *Ponto) *Erro {
	if ponto.Usuario.ID == 0 {
		return &Erro{Codigo: "PONTO10", Mensagem: "Erro ao registra ponto, identificador do usuário inválido.", Err: nil}
	}
	if dt, err := now.Parse(ponto.Data); err != nil {
		return &Erro{Codigo: "PONTO20", Mensagem: "Erro ao realizar parse data ponto.", Err: nil}
	} else if dt.Before(ontem(time.Now())) || dt.After(amanha(time.Now())) {
		return &Erro{Codigo: "PONTO30", Mensagem: `Proibido registrar ponto em uma data posterior ou anterior a data atual.`, Err: nil}
	}
	return nil
}

// NewPontoAvulso - Cria um novo ponto para registro fora da data atual, somente para administradores.
func NewPontoAvulso(ponto *Ponto) *Erro {
	if ponto.Usuario.ID == 0 {
		return &Erro{Codigo: "PONTO10", Mensagem: "Erro ao registra ponto, identificador do usuário inválido.", Err: nil}
	}
	if _, err := now.Parse(ponto.Data); err != nil {
		return &Erro{Codigo: "PONTO20", Mensagem: "Erro ao realizar parse data ponto.", Err: nil}
	}
	return nil
}

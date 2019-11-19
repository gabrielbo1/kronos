package dominio

import (
	"github.com/jinzhu/now"
)

// StatusAtendimento - Controla a situacao do Atendimento
type StatusAtendimento int

const (
	// Aberto - Inicio de um novo atendimento
	Aberto StatusAtendimento = iota + 1
	// Espera - Atendimento iterrompido por alguma situacao
	// de espera de retorno por parte do requerinte.
	Espera
	// Fechado - Atendimento fechado
	Fechado
)

// Intervalo - Data de inicio de um atendimento e data fim.
type Intervalo struct {
	DataInicio string `json:"dataInicio"`
	DataFim    string `json:"dataFim"`
}

// Atendimento - Define unidade do atendo por parte do consultor.
// Teremos a soma da horas de antedimento com as diferencas entre
// as datas de inicio e fim de cada intervalo. Esse controle e necessario
// para atender os casos de atendimentos que prescisam de algum retorno
// por parte do cliente.
type Atendimento struct {
	ID                  int               `json:"id"`
	Usuario             Usuario           `json:"usuario"`
	Cliente             Empresa           `json:"cliente"`
	HorariosAtendimento []Intervalo       `json:"horariosAtendimento"`
	StatusAtendimento   StatusAtendimento `json:"statusAtendimento"`
	Observacao          string            `json:"observacao"`
}

// NewAtendimento - Cria novo atendimento valido.
func NewAtendimento(atnd *Atendimento) *Erro {
	if atnd.Usuario.ID == 0 {
		return &Erro{Codigo: "ATENDIMENTO10", Mensagem: "Erro ao registra atendimento, identificador do consultor inválido."}
	}

	if atnd.Cliente.ID == 0 {
		return &Erro{Codigo: "ATENDIMENTO20", Mensagem: "Erro ao registra atendimento, identificador do cliente inválido."}
	}

	if len(atnd.HorariosAtendimento) == 0 {
		return &Erro{Codigo: "ATENDIMENTO30", Mensagem: "Necessario ao menos a data de inicio para registro do chamado."}
	}

	for i := range atnd.HorariosAtendimento {
		if _, err := now.Parse(atnd.HorariosAtendimento[i].DataInicio); err != nil {
			return &Erro{Codigo: "ATENDIMENTO40", Mensagem: "Erro ao validar data de incio."}
		}

		if atnd.HorariosAtendimento[i].DataFim != "" {
			if _, err := now.Parse(atnd.HorariosAtendimento[i].DataFim); err != nil {
				return &Erro{Codigo: "ATENDIMENTO50", Mensagem: "Erro ao validar data de fim."}
			}
		}

		if atnd.StatusAtendimento == Fechado && atnd.HorariosAtendimento[i].DataFim == "" {
			return &Erro{Codigo: "ATENDIMENTO60", Mensagem: "Para enerrar o atendimento entre com a data fim."}
		}
	}

	return nil
}

package dominio

import (
	"fmt"
	"github.com/jinzhu/now"
	"strconv"
	"strings"
	"time"
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
	ID         int    `json:"id"`
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
		dtInicio, err := now.Parse(atnd.HorariosAtendimento[i].DataInicio)
		if err != nil {
			return &Erro{Codigo: "ATENDIMENTO40", Mensagem: "Erro ao validar data de incio."}
		}
		atnd.HorariosAtendimento[i].DataInicio = dtInicio.Format(time.RFC3339)

		if atnd.HorariosAtendimento[i].DataFim != "" {
			dtFim, err := now.Parse(atnd.HorariosAtendimento[i].DataFim)
			if err != nil || !dtInicio.Before(dtFim) {
				return &Erro{Codigo: "ATENDIMENTO50", Mensagem: "Erro ao validar data de fim, ou data fim antes da data início."}
			}
			atnd.HorariosAtendimento[i].DataFim = dtFim.Format(time.RFC3339)
		}

		if atnd.StatusAtendimento == Fechado && atnd.HorariosAtendimento[i].DataFim == "" {
			return &Erro{Codigo: "ATENDIMENTO60", Mensagem: "Para enerrar o atendimento entre com a data fim."}
		}
	}

	for i := 1; i < len(atnd.HorariosAtendimento); i++ {
		dtFimAnt, _ := now.Parse(atnd.HorariosAtendimento[i-1].DataFim)
		dtInicioAtual, _ := now.Parse(atnd.HorariosAtendimento[i].DataInicio)
		if replaceDateCompre(dtFimAnt.Format(time.RFC3339), dtInicioAtual.Format(time.RFC3339)) {
			msg := "Erro ao atualizar atendimento, novo registro de horário anterior ao horário já cadastrado de início " + dtFimAnt.Format(time.RFC822)
			msg += "."
			msg += " Histórico de horas já lançadas: "
			for k := 0; k < len(atnd.HorariosAtendimento)-1; k++ {
				dtInc, _ := now.Parse(atnd.HorariosAtendimento[i].DataInicio)
				dtFim, _ := now.Parse(atnd.HorariosAtendimento[i].DataFim)
				msg += fmt.Sprintf(" %s -> %s; ", timeFomartPtBr(dtInc), timeFomartPtBr(dtFim))
			}
			return &Erro{Codigo: "ATENDIMENTO70", Mensagem: msg}
		}
	}

	return nil
}

func replaceDateCompre(dt1, dt2 string) bool {
	dt1Num, _ := strconv.Atoi(replaceT((replaceZ(dt1))))
	dt2Num, _ := strconv.Atoi(replaceT((replaceZ(dt2))))
	return dt2Num <= dt1Num
}

func replaceT(dt1 string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(dt1, "T", ""), ":", ""), "-", "")
}

func replaceZ(dt1 string) string {
	return strings.ReplaceAll(strings.ReplaceAll(dt1, ":00Z", ""), ":00-03:00", "")
}

func timeFomartPtBr(date time.Time) string {
	return fmt.Sprintf("%d/%d/%d %d:%d",
		date.Day(),
		date.Month(),
		date.Year(),
		date.Hour(),
		date.Minute())
}

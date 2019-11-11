package visao

import (
	"net/http"

	"github.com/gabrielbo1/kronos/aplicacao"
	"github.com/gabrielbo1/kronos/dominio"
)

// PostAtendimento - POST Atendimento.
func PostAtendimento(w http.ResponseWriter, r *http.Request) {
	var atendimento dominio.Atendimento
	onParserAtendimento(w, readBody(r), atendimento, parseJSON, aplicacao.CadastrarAtendimento)
}

// PutAtendimento - PUT Atendimento.
func PutAtendimento(w http.ResponseWriter, r *http.Request) {
	var atendimento dominio.Atendimento
	onParserAtendimento(w, readBody(r), atendimento, parseJSON, aplicacao.AtualizarAtendimento)
}

// DeleteAtendimento - DELETE Atendimento.
func DeleteAtendimento(w http.ResponseWriter, r *http.Request) {
	var errDominio *dominio.Erro
	var paramInt map[string]int

	if paramInt, errDominio = findURLParamInt(r, []string{"id", "Id atendimento nao passao, erro ao deletar atendimento."}); errDominio == nil {
		if errDominio = aplicacao.ApagarAtendimento(&dominio.Atendimento{ID: paramInt["id"]}); errDominio == nil {
			respostaJSON(w, http.StatusOK, respostaPadraoSimples{Mensagem: "Atendimento apagado com sucesso!"})
			return
		}
	}
	respostaJSON(w, http.StatusBadRequest, errDominio)
}

func onParserAtendimento(w http.ResponseWriter, body []byte, entidade dominio.Atendimento,
	parseFunc func(w http.ResponseWriter, body []byte, entidade interface{}) *dominio.Erro,
	service func(entidade2 *dominio.Atendimento) *dominio.Erro) {
	var e *dominio.Erro
	if e = parseFunc(w, body, &entidade); e == nil {
		if e := service(&entidade); e == nil {
			respostaJSON(w, http.StatusOK, entidade)
		} else {
			respostaJSON(w, http.StatusBadRequest, e)
		}
	}
}

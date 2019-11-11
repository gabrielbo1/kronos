package visao

import (
	"net/http"

	"github.com/gabrielbo1/kronos/aplicacao"
	"github.com/gabrielbo1/kronos/dominio"
)

// PostPonto - POST Ponto.
func PostPonto(w http.ResponseWriter, r *http.Request) {
	var atendimento dominio.Ponto
	onParserPonto(w, readBody(r), atendimento, parseJSON, aplicacao.CadastrarPonto)
}

// PutPonto - PUT Ponto.
func PutPonto(w http.ResponseWriter, r *http.Request) {
	var atendimento dominio.Ponto
	onParserPonto(w, readBody(r), atendimento, parseJSON, aplicacao.AtualizarPonto)
}

// DeletePonto - DELETE Ponto.
func DeletePonto(w http.ResponseWriter, r *http.Request) {
	var errDominio *dominio.Erro
	var paramInt map[string]int

	if paramInt, errDominio = findURLParamInt(r, []string{"id", "Id atendimento nao passao, erro ao deletar atendimento."}); errDominio == nil {
		if errDominio = aplicacao.ApagarPonto(&dominio.Ponto{ID: paramInt["id"]}); errDominio == nil {
			respostaJSON(w, http.StatusOK, respostaPadraoSimples{Mensagem: "Ponto apagado com sucesso!"})
			return
		}
	}
	respostaJSON(w, http.StatusBadRequest, errDominio)
}

func onParserPonto(w http.ResponseWriter, body []byte, entidade dominio.Ponto,
	parseFunc func(w http.ResponseWriter, body []byte, entidade interface{}) *dominio.Erro,
	service func(entidade2 *dominio.Ponto) *dominio.Erro) {
	var e *dominio.Erro
	if e = parseFunc(w, body, &entidade); e == nil {
		if e := service(&entidade); e == nil {
			respostaJSON(w, http.StatusOK, entidade)
		} else {
			respostaJSON(w, http.StatusBadRequest, e)
		}
	}
}

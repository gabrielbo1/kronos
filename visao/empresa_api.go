package visao

import (
	"net/http"

	"github.com/gabrielbo1/kronos/aplicacao"
	"github.com/gabrielbo1/kronos/dominio"
)

// PostEmpresa - POST Empresa.
func PostEmpresa(w http.ResponseWriter, r *http.Request) {
	var atendimento dominio.Empresa
	onParserEmpresa(w, readBody(r), atendimento, parseJSON, aplicacao.CadastrarEmpresa)
}

// PutEmpresa - PUT Empresa.
func PutEmpresa(w http.ResponseWriter, r *http.Request) {
	var atendimento dominio.Empresa
	onParserEmpresa(w, readBody(r), atendimento, parseJSON, aplicacao.AtualizarEmpresa)
}

// DeleteEmpresa - DELETE Empresa.
func DeleteEmpresa(w http.ResponseWriter, r *http.Request) {
	var errDominio *dominio.Erro
	var paramInt map[string]int

	if paramInt, errDominio = findURLParamInt(r, []string{"id", "Id atendimento nao passao, erro ao deletar atendimento."}); errDominio == nil {
		if errDominio = aplicacao.ApagarEmpresa(&dominio.Empresa{ID: paramInt["id"]}); errDominio == nil {
			respostaJSON(w, http.StatusOK, respostaPadraoSimples{Mensagem: "Empresa apagado com sucesso!"})
			return
		}
	}
	respostaJSON(w, http.StatusBadRequest, errDominio)
}

// GetEmpresas - GET Empresas.
func GetEmpresas(w http.ResponseWriter, r *http.Request) {
	empresas, e := aplicacao.BuscaEmpresas()
	if e == nil {
		respostaJSON(w, http.StatusOK, empresas)
		return
	}
	respostaJSON(w, http.StatusBadRequest, e)
}

func onParserEmpresa(w http.ResponseWriter, body []byte, entidade dominio.Empresa,
	parseFunc func(w http.ResponseWriter, body []byte, entidade interface{}) *dominio.Erro,
	service func(entidade2 *dominio.Empresa) *dominio.Erro) {
	var e *dominio.Erro
	if e = parseFunc(w, body, &entidade); e == nil {
		if e := service(&entidade); e == nil {
			respostaJSON(w, http.StatusOK, entidade)
		} else {
			respostaJSON(w, http.StatusBadRequest, e)
		}
	}
}

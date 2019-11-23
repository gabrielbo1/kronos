package visao

import (
	"net/http"
	"strconv"

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

	if paramInt, errDominio = findURLParamInt(r, []string{"id", "Id atendimento não foi passado, erro ao deletar atendimento."}); errDominio == nil {
		if errDominio = aplicacao.ApagarAtendimento(&dominio.Atendimento{ID: paramInt["id"]}); errDominio == nil {
			respostaJSON(w, http.StatusOK, respostaPadraoSimples{Mensagem: "Atendimento apagado com sucesso!"})
			return
		}
	}
	respostaJSON(w, http.StatusBadRequest, errDominio)
}

// GetAtendimento - GET Atendimentos
func GetAtendimento(w http.ResponseWriter, r *http.Request) {
	var errDominio *dominio.Erro
	var paramInt map[string]int

	if paramInt, errDominio = findURLParamInt(r, []string{"id", "Id usuário não foi passado, erro ao deletar atendimento."}); errDominio == nil {
		if atendimentos, errDominio := aplicacao.BuscarAtendimentoIdUsuario(paramInt["id"]); errDominio == nil {
			respostaJSON(w, http.StatusOK, atendimentos)
			return
		}
	}
	respostaJSON(w, http.StatusBadRequest, errDominio)
}

// GetAtendimentoPaginado - GET Atendimentos busca paginada.
func GetAtendimentoPaginado(w http.ResponseWriter, r *http.Request) {
	var errDominio *dominio.Erro
	var paramInt map[string]int
	var pagina dominio.Pagina

	if paramInt, errDominio = findURLParamInt(r, []string{
		"id", "Id usuário não foi passado, erro ao buscar atendimentos.",
		"numPag", "Número da página não foi passado, erro ao buscar atendimentos.",
		"qtdPag", "Quantidade  por página não foi passado, erro ao buscar atendimentos.",
	}); errDominio == nil {
		if pagina, errDominio = aplicacao.BuscarAtendimentoIdUsuarioPaginado(paramInt["id"],
			dominio.Pagina{NumPagina: paramInt["numPag"], QtdPorPagina: paramInt["qtdPag"]}); errDominio == nil {
			respostaJSON(w, http.StatusOK, pagina)
			return
		}
	}
	respostaJSON(w, http.StatusBadRequest, errDominio)
}

// GetAtendimentoPaginadoLike - Busca paginada de atendimento com like.
func GetAtendimentoPaginadoLike(w http.ResponseWriter, r *http.Request) {
	var errDominio *dominio.Erro
	var paramUrl map[string]string
	var pagina dominio.Pagina

	if paramUrl, errDominio = findURLParam(r, []string{
		"id", "Id usuário não foi passado, erro ao buscar atendimentos.",
		"numPag", "Número da página não foi passado, erro ao buscar atendimentos.",
		"qtdPag", "Quantidade  por página não foi passado, erro ao buscar atendimentos.",
		"like", "Texto para busca não foi passado, erro ao buscar atendimentos.",
	}); errDominio == nil {
		id, errId := strconv.Atoi(paramUrl["id"])
		numPag, errNumPag := strconv.Atoi(paramUrl["numPag"])
		qtdPag, errQtdPag := strconv.Atoi(paramUrl["qtdPag"])
		if errId != nil || errNumPag != nil || errQtdPag != nil {
			errDominio = &dominio.Erro{Codigo: "", Mensagem: `Erro ao havaliar parâmetros get, id, número da página 
														    e quantidade  por página do tipo inteiro.`}
		} else {
			if pagina, errDominio = aplicacao.BuscarAtendimentoIdUsuarioLikePaginado(id,
				paramUrl["like"], dominio.Pagina{NumPagina: numPag, QtdPorPagina: qtdPag}); errDominio == nil {
				respostaJSON(w, http.StatusOK, pagina)
				return
			}
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

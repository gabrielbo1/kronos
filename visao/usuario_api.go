package visao

import (
	"net/http"

	"github.com/gabrielbo1/kronos/aplicacao"
	"github.com/gabrielbo1/kronos/dominio"
)

// PostUsuario - POST Usuario.
func PostUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario dominio.Usuario
	onParserUsuario(w, readBody(r), usuario, parseJSON, aplicacao.CadastrarUsuario)
}

// PutUsuario - PUT Usuario.
func PutUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario dominio.Usuario
	onParserUsuario(w, readBody(r), usuario, parseJSON, aplicacao.AtualizarUsuario)
}

// GetUsuarios - GET Usuarios.
func GetUsuarios(w http.ResponseWriter, r *http.Request) {
	usuarios, e := aplicacao.BuscarUsuarios()
	if e == nil {
		respostaJSON(w, http.StatusOK, usuarios)
		return
	}
	respostaJSON(w, http.StatusBadRequest, e)
}

// PostUsuario - Login usuario.
func PostLoginUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario dominio.Usuario
	var e *dominio.Erro
	if e = parseJSON(w, readBody(r), &usuario); e == nil {
		if usuario, e = aplicacao.Login(usuario.Login, usuario.Senha); e == nil {
			respostaJSON(w, http.StatusOK, usuario)
			return
		}
	}
	respostaJSON(w, http.StatusBadRequest, e)
}

// GetUsuarioOk - GET Confir token e login.
func GetUsuarioOk(w http.ResponseWriter, r *http.Request) {
	respostaJSON(w, http.StatusOK, respostaPadraoSimples{Mensagem: "OK"})
}

func onParserUsuario(w http.ResponseWriter, body []byte, entidade dominio.Usuario,
	parseFunc func(w http.ResponseWriter, body []byte, entidade interface{}) *dominio.Erro,
	service func(entidade2 *dominio.Usuario) *dominio.Erro) {
	var e *dominio.Erro
	if e = parseFunc(w, body, &entidade); e == nil {
		if e := service(&entidade); e == nil {
			respostaJSON(w, http.StatusOK, entidade)
		} else {
			respostaJSON(w, http.StatusBadRequest, e)
		}
	}
}

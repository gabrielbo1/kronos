package visao

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var rotas = Routes{
	Route{
		Name:        "PostAtendimento",
		Method:      "POST",
		Pattern:     "/atendimento",
		HandlerFunc: PostAtendimento,
	},
	Route{
		Name:        "PutAtendimento",
		Method:      "PUT",
		Pattern:     "/atendimento",
		HandlerFunc: PutAtendimento,
	},
	Route{
		Name:        "GetAtendimento",
		Method:      "GET",
		Pattern:     "/atendimento/usuario/{id}",
		HandlerFunc: GetAtendimento,
	},
	Route{
		Name:        "GetAtendimentoPaginado",
		Method:      "GET",
		Pattern:     "/atendimento/usuario/{id}/{numPag}/{qtdPag}",
		HandlerFunc: GetAtendimentoPaginado,
	},
	Route{
		Name:        "GetAtendimentoPaginadoLike",
		Method:      "GET",
		Pattern:     "/atendimento/usuario/{id}/{numPag}/{qtdPag}/{like}",
		HandlerFunc: GetAtendimentoPaginadoLike,
	},
	Route{
		Name:        "DeleteAtendimento",
		Method:      "DELETE",
		Pattern:     "/atendimento/{id}",
		HandlerFunc: DeleteAtendimento,
	},
	Route{
		Name:        "PostEmpresa",
		Method:      "POST",
		Pattern:     "/empresa",
		HandlerFunc: PostEmpresa,
	},
	Route{
		Name:        "PutEmpresa",
		Method:      "PUT",
		Pattern:     "/empresa",
		HandlerFunc: PutEmpresa,
	},
	Route{
		Name:        "DeleteEmpresa",
		Method:      "DELETE",
		Pattern:     "/empresa/{id}",
		HandlerFunc: DeleteEmpresa,
	},
	Route{
		Name:        "GetEmpresas",
		Method:      "GET",
		Pattern:     "/empresa",
		HandlerFunc: GetEmpresas,
	},
	Route{
		Name:        "PostPonto",
		Method:      "POST",
		Pattern:     "/ponto",
		HandlerFunc: PostPonto,
	},
	Route{
		Name:        "PutPonto",
		Method:      "PUT",
		Pattern:     "/ponto",
		HandlerFunc: PutPonto,
	},
	Route{
		Name:        "DeletePonto",
		Method:      "DELETE",
		Pattern:     "/ponto/{id}",
		HandlerFunc: DeletePonto,
	},
	Route{
		Name:        "PostUsuario",
		Method:      "POST",
		Pattern:     "/usuario",
		HandlerFunc: PostUsuario,
	},
	Route{
		Name:        "PutUsuario",
		Method:      "PUT",
		Pattern:     "/usuario",
		HandlerFunc: PutUsuario,
	},
	Route{
		Name:        "GetUsuarios",
		Method:      "GET",
		Pattern:     "/usuario",
		HandlerFunc: GetUsuarios,
	},
	Route{
		Name:        "PostLoginUsuario",
		Method:      "POST",
		Pattern:     "/usuario/login",
		HandlerFunc: PostLoginUsuario,
	},
	Route{
		Name:        "GetUsuarioOk",
		Method:      "GET",
		Pattern:     "/usuario/loginok",
		HandlerFunc: GetUsuarioOk,
	},
}

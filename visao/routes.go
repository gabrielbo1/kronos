package visao

import "net/http"

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routes []route

var rotas = routes{
	route{
		Name:        "PostAtendimento",
		Method:      "POST",
		Pattern:     "/atendimento",
		HandlerFunc: PostAtendimento,
	},
	route{
		Name:        "PutAtendimento",
		Method:      "PUT",
		Pattern:     "/atendimento",
		HandlerFunc: PutAtendimento,
	},
	route{
		Name:        "DeleteAtendimento",
		Method:      "DELETE",
		Pattern:     "/atendimento",
		HandlerFunc: DeleteAtendimento,
	},
	route{
		Name:        "PostEmpresa",
		Method:      "POST",
		Pattern:     "/empresa",
		HandlerFunc: PostEmpresa,
	},
	route{
		Name:        "PutEmpresa",
		Method:      "PUT",
		Pattern:     "/empresa",
		HandlerFunc: PutEmpresa,
	},
	route{
		Name:        "DeleteEmpresa",
		Method:      "DELETE",
		Pattern:     "/empresa",
		HandlerFunc: DeleteEmpresa,
	},
	route{
		Name:        "GetEmpresas",
		Method:      "GET",
		Pattern:     "/empresa",
		HandlerFunc: GetEmpresas,
	},
	route{
		Name:        "PostPonto",
		Method:      "POST",
		Pattern:     "/ponto",
		HandlerFunc: PostPonto,
	},
	route{
		Name:        "PutPonto",
		Method:      "PUT",
		Pattern:     "/ponto",
		HandlerFunc: PutPonto,
	},
	route{
		Name:        "DeletePonto",
		Method:      "DELETE",
		Pattern:     "/ponto",
		HandlerFunc: DeletePonto,
	},
	route{
		Name:        "PostUsuario",
		Method:      "POST",
		Pattern:     "/usuario",
		HandlerFunc: PostUsuario,
	},
	route{
		Name:        "PutUsuario",
		Method:      "PUT",
		Pattern:     "/usuario",
		HandlerFunc: PutUsuario,
	},
	route{
		Name:        "GetUsuarios",
		Method:      "GET",
		Pattern:     "/usuario",
		HandlerFunc: GetUsuarios,
	},
	route{
		Name:        "PostLoginUsuario",
		Method:      "POST",
		Pattern:     "/usuario/login",
		HandlerFunc: PostLoginUsuario,
	},
}

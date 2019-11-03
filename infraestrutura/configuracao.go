package infraestrutura

// Configuracao - Define parametros de configuracao da aplicao.
type Configuracao struct {
	LoginAdm         string `json:"loginAdmin"`
	SenhaAdm         string `json:"senhaAdmin"`
	Banco            string `json:"banco"`
	IPBanco          string `json:"ipBanco"`
	NomeBanco        string `json:"nomeBanco"`
	PortaBanco       int    `json:"portaBanco"`
	UsuarioBanco     string `json:"usuarioBanco"`
	SenhaBanco       string `json:"senhaBanco"`
	DiretorioScripts string `json:"diretorioScripts"`
}

//Config - Configuracao default
var Config Configuracao = Configuracao{
	LoginAdm:         "admin",
	SenhaAdm:         "adm123",
	Banco:            "POSTGRES",
	IPBanco:          "localhost",
	NomeBanco:        "kronos",
	PortaBanco:       5432,
	UsuarioBanco:     "kronospostgres",
	SenhaBanco:       "123456",
	DiretorioScripts: "",
}

package infraestrutura

import (
	"os"
	"strconv"
	"time"
)

// Configuracao - Define parâmetros de configuracao da aplicao.
type Configuracao struct {
	LoginAdm         string `json:"loginAdmin"`
	SenhaAdm         string `json:"senhaAdmin"`
	Banco            string `json:"banco"`
	IPBanco          string `json:"ipBanco"`
	NomeBanco        string `json:"nomeBanco"`
	PortaBanco       int    `json:"portaBanco"`
	UsuarioBanco     string `json:"usuarioBanco"`
	SenhaBanco       string `json:"senhaBanco"`
	SslBanco         string `json:"sslBanco"`
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
	SslBanco:         "disable",
	DiretorioScripts: "./infraestrutura/repositorio/script_postgressql",
}

// ConfigInit - Chega parametros e ajusta configuracao
func ConfigInit() {
	if p := os.Getenv("HOST_POSTGRES"); p != "" {
		Config.IPBanco = p
	}
	if p := os.Getenv("PORTA_POSTGRES"); p != "" {
		Config.PortaBanco, _ = strconv.Atoi(p)
	}
	if p := os.Getenv("BANCO_POSTGRES"); p != "" {
		Config.NomeBanco = p
	}
	if p := os.Getenv("USUARIO_POSTGRES"); p != "" {
		Config.UsuarioBanco = p
	}
	if p := os.Getenv("SENHA_POSTGRES"); p != "" {
		Config.SenhaBanco = p
	}
	if p := os.Getenv("SSL_BANCO"); p != "" {
		Config.SslBanco = p
	}
}

func TimeZonePadrao() *time.Location {
	localeApp, _ := time.LoadLocation("America/Sao_Paulo")
	return localeApp
}

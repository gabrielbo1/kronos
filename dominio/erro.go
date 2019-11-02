package dominio

import "fmt"

// Erro - Define formato padrão de retorno de mensagem da aplicação.
type Erro struct {
	Codigo   string `json:"codigo"`
	Mensagem string `json:"mensagem"`
	Err      error  `json:"err"`
}

// Erro - Serializa erro em string.
func (e Erro) Erro() string {
	return fmt.Sprintf("%s - %s", e.Codigo, e.Mensagem)
}

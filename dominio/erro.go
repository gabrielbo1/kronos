package dominio

import (
	"fmt"
	"github.com/pkg/errors"
)

// Erro - Define formato padrão de retorno de mensagem da aplicação.
type Erro struct {
	Codigo   string `json:"codigo"`
	Mensagem string `json:"mensagem"`
	Err      error  `json:"err"`
}

func OnError(e *Erro) error {
	if e == nil {
		return nil
	}
	return errors.New(fmt.Sprintf("%s - %s", e.Codigo, e.Mensagem))
}

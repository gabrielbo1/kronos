package dominio

// Empresa - Define entidade empresa, cliente dos chamdos.
type Empresa struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Ativa bool   `json:"ativa"`
}

// NewEmpresa - Cria nova empresa válida.
func NewEmpresa(empresa *Empresa) *Erro {
	if empresa.Nome == "" {
		return &Erro{Codigo: "EMPRESA10", Mensagem: "Erro ao criar nova empresa campo nome obrigatório", Err: nil}
	}
	empresa.Ativa = true
	return nil
}

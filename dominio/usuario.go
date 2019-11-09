package dominio

// Rotina - Rotina da interface com usuário.
type Rotina struct {
	ID     int    `json:"id"`
	Rotina string `json:"rotina"`
}

// NewRotina - Valida rotina.
func NewRotina(rotina *Rotina) *Erro {
	if rotina.Rotina == "" {
		return &Erro{Codigo: "ROTINA10", Mensagem: "Nome rotina vazio, inválido para cadastro.", Err: nil}
	}
	return nil
}

// Acesso - Define nível de acesso de uma rotina
// controle de perimissões de usuario para cada rotina
// do sistema.
type Acesso struct {
	Rotina     Rotina `json:"rotina"`
	Criar      bool   `json:"criar"`
	Atualizar  bool   `json:"atualizar"`
	Deletar    bool   `json:"deletar"`
	Visualizar bool   `json:"visualizar"`
}

// NewAcesso - Valida acesso.
func NewAcesso(acesso *Acesso) *Erro {
	if e := NewRotina(&acesso.Rotina); e != nil {
		return e
	}

	if !acesso.Criar && !acesso.Atualizar && !acesso.Deletar && !acesso.Visualizar {
		return &Erro{Codigo: "ACESSO10", Mensagem: `Erro ao criar acesso, 
			para criar acesso necessário ao menos umas permissões
			(Criar, Atualizar, Deletar ou Visualizar)`, Err: nil}
	}

	return nil
}

// Usuario - Define entidade usuário com respecitivo nível de acesso.
type Usuario struct {
	ID     int      `json:"id"`
	Nome   string   `json:"nome"`
	Login  string   `json:"login"`
	Senha  string   `json:"senha"`
	Acesso []Acesso `json:"acesso"`
}

// NewUsuario - Valida novo usuario.
func NewUsuario(usuario *Usuario) *Erro {
	if usuario.Nome == "" {
		return &Erro{Codigo: "USUARIO10", Mensagem: "Erro ao criar novo usuário, campo nome obrigatório.", Err: nil}
	}

	if usuario.Login == "" {
		return &Erro{Codigo: "USUARIO20", Mensagem: "Erro ao criar novo usuário, campo login obrigatório.", Err: nil}
	}

	if usuario.Senha == "" {
		return &Erro{Codigo: "USUARIO30", Mensagem: "Erro ao criar novo usuário, campo senha obrigatório.", Err: nil}
	}

	if len(usuario.Acesso) == 0 {
		return &Erro{Codigo: "USUARIO40", Mensagem: "Erro ao criar novo usuário, necessário ao menos uma rotina associada.", Err: nil}
	}

	for i := range usuario.Acesso {
		if e := NewAcesso(&usuario.Acesso[i]); e != nil {
			return e
		}
	}

	return nil
}

package aplicacao

import (
	"database/sql"

	"github.com/gabrielbo1/kronos/dominio"
	"github.com/gabrielbo1/kronos/infraestrutura/repositorio"
)

var repUsuario = repositorio.NewUsuarioRepositorio()

// CadastrarUsuario - Cadastro de usuario.
func CadastrarUsuario(usuario *dominio.Usuario) (errDominio *dominio.Erro) {
	if errDominio = dominio.NewUsuario(usuario); errDominio != nil {
		return errDominio
	}

	if errTX := repositorio.Transact(repositorio.DB, func(tx *sql.Tx) error {
		if usuario, err := repUsuario.BuscaLogin(tx, usuario.Login); err != nil || usuario.ID != 0 {
			if err != nil {
				errDominio = err
				return nil
			}

			if usuario.ID != 0 {
				errDominio = &dominio.Erro{Codigo: "USUARIO_SERVICO10", Mensagem: "Login já cadastrado."}
				return nil
			}
		}
		usuario.ID, errDominio = repUsuario.Save(tx, *usuario)
		return nil
	}); errTX != nil {
		return TrataErroConexao(errDominio, errTX)
	}

	return errDominio
}

// AtualizarUsuario - Atualizar usuario.
func AtualizarUsuario(usuario *dominio.Usuario) (errDominio *dominio.Erro) {
	if errDominio = dominio.NewUsuario(usuario); errDominio != nil {
		return errDominio
	}

	if errTX := repositorio.Transact(repositorio.DB, func(tx *sql.Tx) error {
		if usu, err := repUsuario.BuscaLogin(tx, usuario.Login); err != nil || (usu.ID != 0 && usu.ID != usuario.ID) {
			if err != nil {
				errDominio = err
				return nil
			}

			if usu.ID != 0 && usu.ID != usuario.ID {
				errDominio = &dominio.Erro{Codigo: "USUARIO_SERVICO10", Mensagem: "Login já cadastrado."}
				return nil
			}
		}
		errDominio = repUsuario.Update(tx, *usuario)
		return nil
	}); errTX != nil {
		return TrataErroConexao(errDominio, errTX)
	}
	return errDominio
}

// BuscarUsuarios - Buscar todos usuarios.
func BuscarUsuarios() (usuarios []dominio.Usuario, errDominio *dominio.Erro) {
	if errTX := repositorio.Transact(repositorio.DB, func(tx *sql.Tx) error {
		usuarios, errDominio = repUsuario.FindAll(tx)
		return nil
	}); errTX != nil {
		return usuarios, TrataErroConexao(errDominio, errTX)
	}
	return usuarios, errDominio
}

// Login - Login usuario e senha.
func Login(login, senha string) (usuario dominio.Usuario, errDominio *dominio.Erro) {
	if errTX := repositorio.Transact(repositorio.DB, func(tx *sql.Tx) error {
		usuario, errDominio = repUsuario.Login(tx, login, senha)
		return nil
	}); errTX != nil {
		return usuario, TrataErroConexao(errDominio, errTX)
	}
	return usuario, errDominio
}

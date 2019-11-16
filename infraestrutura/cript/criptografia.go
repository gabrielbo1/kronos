package cript

import "github.com/gabrielbo1/kronos/dominio"

type Criptografia interface {
	Criptografar(bytesArray []byte) (erro *dominio.Erro, bytesCriptografados []byte)
	Decriptografar(bytesArray []byte) (erro *dominio.Erro, bytesDecriptografados []byte)
}

// FabricaCriptografia - Instancia algoritomo "aes" ou "rsa"
func FabricaCriptografia(algoritmo string) Criptografia {
	switch algoritmo {
	case "aes":
		return newAES(&algoritmoAes{})
	case "rsa":
		return newRSA(&algoritmoRsa{})
	}
	return nil
}

// RSA OATUH
var RsaToken Criptografia = FabricaCriptografia("rsa")

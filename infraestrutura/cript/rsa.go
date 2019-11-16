package cript

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"log"

	"github.com/gabrielbo1/kronos/dominio"
)

type algoritmoRsa struct {
	chavePrivada *rsa.PrivateKey
	chavePublica *rsa.PublicKey
}

func newRSA(rsaThis *algoritmoRsa) *algoritmoRsa {
	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	rsaThis.chavePrivada = rsaKey
	rsaThis.chavePublica = &rsaKey.PublicKey
	return rsaThis
}

// Criptografar - criptografa sequecencia de bytes.
func (rsaThis *algoritmoRsa) Criptografar(bytesArray []byte) (erro *dominio.Erro, bytesCriptografados []byte) {
	if len(bytesArray) > 190 {
		return &dominio.Erro{"RSA-10", "Erro ao criptografar algoritmo RSA.", nil}, nil

	}
	bytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, rsaThis.chavePublica, bytesArray, []byte(""))
	if err != nil {
		log.Println(err.Error())
		return &dominio.Erro{"RSA-10", "Erro ao criptografar algoritmo RSA.", err}, nil
	}
	return nil, bytes
}

// Decriptografar - decriptograda a sequecia de bytes.
func (rsaThis *algoritmoRsa) Decriptografar(bytesArray []byte) (erro *dominio.Erro, bytesDecriptografados []byte) {
	bytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, rsaThis.chavePrivada, bytesArray, []byte(""))
	if err != nil {
		log.Println(err.Error())
		return &dominio.Erro{"RSA-20", "Erro ao decriptografar algoritmo RSA.", err}, nil
	}
	return nil, bytes
}

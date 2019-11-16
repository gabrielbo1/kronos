package cript

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"log"

	"github.com/gabrielbo1/kronos/dominio"
)

type algoritmoAes struct {
	keyBytes             []byte
	vetorDeInicializacao []byte
}

func newAES(aes *algoritmoAes) *algoritmoAes {
	keyString := newStringNumberRand(64)
	aes.keyBytes, _ = hex.DecodeString(keyString)
	aes.vetorDeInicializacao = make([]byte, 12)
	rand.Read(aes.vetorDeInicializacao)
	return aes
}

// Criptografar - criptografa usando AES.
func (aesThis *algoritmoAes) Criptografar(bytesArray []byte) (erro *dominio.Erro, bytesCriptografados []byte) {
	block, err := aes.NewCipher(aesThis.keyBytes)
	if err != nil {
		log.Println(err.Error())
		return &dominio.Erro{"AES-10", "Erro ao criptografar algoritmo AES.", err}, nil
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println(err.Error())
		return &dominio.Erro{"AES-10", "Erro ao criptografar algoritmo AES.", err}, nil
	}
	return nil, aesgcm.Seal(nil, aesThis.vetorDeInicializacao, bytesArray, nil)
}

// Criptografar - descriptografa usando AES.
func (aesThis *algoritmoAes) Decriptografar(bytesArray []byte) (erro *dominio.Erro, bytesDecriptografados []byte) {
	block, err := aes.NewCipher(aesThis.keyBytes)
	if err != nil {
		log.Println(err.Error())
		return &dominio.Erro{"AES-20", "Erro ao decriptografar algoritmo AES.", err}, nil
	}
	aesgcm, err := cipher.NewGCM(block)
	bytesDecritogrados, err := aesgcm.Open(nil, aesThis.vetorDeInicializacao, bytesArray, nil)
	if err != nil {
		log.Println(err.Error())
		return &dominio.Erro{"AES-20", "Erro ao decriptografar algoritmo AES.", err}, nil
	}
	return nil, bytesDecritogrados
}

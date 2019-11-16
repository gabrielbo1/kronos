package cript

import (
	"fmt"
	"testing"
)

func TestAES(t *testing.T) {
	var cript Criptografia = FabricaCriptografia("aes")
	s := "TEXTO QUALQUER"
	_, arrayCript := cript.Criptografar([]byte(s))
	_, arrayDecript := cript.Decriptografar(arrayCript)
	if s != string(arrayDecript) {
		t.Fail()
	}
	fmt.Printf("%s, %s\n", string(arrayCript), string(arrayDecript))
	cript = FabricaCriptografia("aes")
}

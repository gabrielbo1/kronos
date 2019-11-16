package cript

import (
	"fmt"
	"testing"
)

func TestRSA(t *testing.T) {
	var cript Criptografia = FabricaCriptografia("rsa")
	s := "TEXTOQUALQUERTEXTOQUALQUERTEXTOQUALQUERTEXTOQUALQUERTEXTOQUALQUERTEXTOQUALQUERTEXTOQUALQUERTEXTOQUALQUERTEXTOQUALQUERTEXTOQUALQUERTEXTOQUALQUERTEXTOQUALQUERTEXTOQUALQUERTEXTOQUALQUERTEXTOQUA"
	_, arrayCript := cript.Criptografar([]byte(s))
	_, arrayDecript := cript.Decriptografar(arrayCript)
	if s != string(arrayDecript) {
		t.Fail()
	}
	fmt.Printf("%s\n %s\n", string(arrayCript), string(arrayDecript))
	cript = FabricaCriptografia("aes")
}

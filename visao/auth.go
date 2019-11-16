package visao

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/gabrielbo1/kronos/aplicacao"
	"github.com/gabrielbo1/kronos/infraestrutura/cript"
)

var limparCacherStart = false
var cache map[string]bool = make(map[string]bool)

func basicAuth(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !limparCacherStart {
			go limparCache()
			limparCacherStart = true
		}

		username, token, authOK := r.BasicAuth()
		if authOK == false {
			http.Error(w, "Not authorized", 401)
			return
		}

		if !validToken(username, token) {
			http.Error(w, "Not authorized", 401)
			return
		}

		inner.ServeHTTP(w, r)
	})
}

func validToken(username string, token string) bool {
	if val, ok := cache[token]; ok {
		return val
	}

	if tokenBytes, err := base64.StdEncoding.DecodeString(token); err == nil {
		if err, tokenDecrit := cript.RsaToken.Decriptografar(tokenBytes); err == nil {
			if login, err := aplicacao.Login(username, string(tokenDecrit)); err == nil && login.ID != 0 {
				cache[token] = true
				return true
			}
		}
	}

	cache[token] = false
	return false
}

func limparCache() {
	horas12, _ := time.ParseDuration("12h")
	for {
		time.Sleep(horas12)
		cache = make(map[string]bool)
		cript.RsaToken = cript.FabricaCriptografia("rsa")
	}
}

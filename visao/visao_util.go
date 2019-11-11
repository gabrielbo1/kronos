package visao

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gabrielbo1/kronos/dominio"
	"github.com/gorilla/mux"
)

type respostaPadraoSimples struct {
	Mensagem string `json:"mensagem"`
}

func respostaJSON(w http.ResponseWriter, code int, entidade interface{}) {
	response, _ := json.Marshal(entidade)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func readBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	return body
}

func parseJSON(w http.ResponseWriter, body []byte, entidade interface{}) *dominio.Erro {
	if err := json.Unmarshal(body, entidade); err != nil {
		var erro dominio.Erro = dominio.Erro{"RESPOSTA_10", "ERRO AO REALIZAR DESERIALIZAR JSON", err}
		log.Println(err)
		respostaJSON(w, http.StatusUnprocessableEntity, erro)
		return &erro
	}
	return nil
}

func findURLParam(r *http.Request, keys []string) (keyResolv map[string]string, erro *dominio.Erro) {
	result := make(map[string]string)
	vars := mux.Vars(r)
	for i := 0; i < len(keys); i += 2 {
		if k := vars[keys[i]]; k == "" {
			return nil, &dominio.Erro{"ERRO_PARAM_URL", keys[i+1], nil}
		} else {
			result[keys[i]] = k
		}
	}
	return result, nil
}

func findURLParamInt(r *http.Request, keys []string) (keyResolv map[string]int, erro *dominio.Erro) {
	result := make(map[string]int)
	vars := mux.Vars(r)
	for i := 0; i < len(keys); i += 2 {
		if k, err := strconv.Atoi(vars[keys[i]]); err != nil {
			return nil, &dominio.Erro{"ERRO_PARAM_URL", keys[i+1], err}
		} else {
			result[keys[i]] = k
		}
	}
	return result, nil
}

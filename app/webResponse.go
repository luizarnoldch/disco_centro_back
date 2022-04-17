package app

import (
	"encoding/json"
	"net/http"
)

func writeResponse(w http.ResponseWriter, codigo int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(codigo)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

package tools

import (
	"encoding/json"
	"net/http"
)

// RespondWithError respuesta con un error
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithjson(w, code, map[string]string{"error": msg})
}

//RespondWithjson respuesta con json
func RespondWithjson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

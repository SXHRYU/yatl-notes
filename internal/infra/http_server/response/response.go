package http_server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteErrorResponse(w http.ResponseWriter, text string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintf(w, `{"error":"%s"}`, text)
}

func WriteResponse(w http.ResponseWriter, resp any, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

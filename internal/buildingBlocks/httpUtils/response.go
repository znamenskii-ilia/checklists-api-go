package httpUtils

import (
	"encoding/json"
	"net/http"
)

func DecodeBody(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

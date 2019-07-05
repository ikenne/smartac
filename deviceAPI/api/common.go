package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *APIServer) responseJSON(w http.ResponseWriter, status int, v interface{}) {
	// r, err := json.MarshalIndent(v, "", " ")
	r, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	s.write(w, status, r)
}

func (s *APIServer) write(w http.ResponseWriter, status int, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if _, err := w.Write(body); err != nil {
		fmt.Println("Error writing response")
	}
}

func (s *APIServer) responsError(w http.ResponseWriter, code int, message string) {
	s.responseJSON(w, code, struct {
		Error string `json:"error"`
	}{message})
}

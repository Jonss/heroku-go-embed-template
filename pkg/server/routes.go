package server

import (
	"net/http"
)

func (s *Server) Routes() {
	http.HandleFunc("/api/hello", s.Hello())
}
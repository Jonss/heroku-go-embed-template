package server

import (
	"fmt"
	"net/http"
)

func (s *Server) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, heroku-go-embed-template!")
	}
}
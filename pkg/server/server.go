package server

import (
	"github.com/Jonss/heroku-go-embed-template/pkg/config"
	"github.com/gorilla/mux"
)

type Server struct {
	router            *mux.Router
	config            config.Config
	restValidator     *Validator
}

func NewServer(
	r *mux.Router,
	cfg config.Config,
	v *Validator,
) *Server {
	return &Server{
		router: r,
		config:            cfg,
		restValidator: v,
	}
}
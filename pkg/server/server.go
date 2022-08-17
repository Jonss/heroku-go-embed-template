package server

import (
	"github.com/Jonss/heroku-go-embed-template/pkg/config"
)

type Server struct {
	config            config.Config
	restValidator     *Validator
}

func NewServer(
	cfg config.Config,
	v *Validator,
) *Server {
	return &Server{
		config:            cfg,
		restValidator: v,
	}
}
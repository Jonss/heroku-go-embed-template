package server

func (s *Server) Routes() {
	s.router.HandleFunc("/api/hello", s.Hello())
}
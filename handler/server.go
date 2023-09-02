package handler

import (
	"net/http"

	"testing-nextalent/pkg/grace"
)

// ApiHandler ...
type ApiHandler interface {
	// Masukkan fungsi handler di sini
	ApiHandler(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	server *http.Server
	Api    ApiHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	return grace.Serve(port, s.Handler())
}

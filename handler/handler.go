package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Handler will initialize mux router and register handler
func (s *Server) Handler() *mux.Router {
	r := mux.NewRouter()

	// Tambahan Prefix di depan API endpoint
	router := r.PathPrefix("/api").Subrouter()
	// Health Check
	router.HandleFunc("", defaultHandler).Methods("GET")
	router.HandleFunc("/", defaultHandler).Methods("GET")
	// Routes
	router.HandleFunc("/country", s.Api.ApiHandler).Methods("GET", "POST")

	return r
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Service API"))
}

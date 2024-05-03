package server

import (
	"net/http"
	
	"github.com/gorilla/mux"
)

func router(h *Handler) *mux.Router {
	r := mux.NewRouter()
	r = r.SkipClean(true)
	r.HandleFunc("/ping", h.Ping).Methods(http.MethodGet)
	return r
}

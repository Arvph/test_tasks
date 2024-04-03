package server

import (
	"github.com/gorilla/mux"
)

func routers(conf *Server) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/tasks", conf.Handler.Create).Methods("POST")
	r.HandleFunc("/tasks", conf.Handler.GetAll).Methods("GET")
	r.HandleFunc("/tasks/{id}", conf.Handler.GetByID).Methods("GET")
	r.HandleFunc("/tasks/{id}", conf.Handler.UpdateByID).Methods("PUT")
	r.HandleFunc("/tasks/{id}", conf.Handler.DeleteByID).Methods("DELETE")

	return r
}

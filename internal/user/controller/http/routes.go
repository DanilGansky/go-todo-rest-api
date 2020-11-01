package http

import "github.com/gorilla/mux"

// SetupRoutes ...
func SetupRoutes(router *mux.Router, controller *UserController) {
	router.HandleFunc("/users", controller.GetAll).Methods("GET")
	router.HandleFunc("/users", controller.Post).Methods("POST")
	router.HandleFunc("/users/{id}", controller.GetSingle).Methods("GET")
	router.HandleFunc("/users/{id}", controller.Put).Methods("PUT")
	router.HandleFunc("/users/{id}", controller.Delete).Methods("DELETE")
}

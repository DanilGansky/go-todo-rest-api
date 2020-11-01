package http

import "github.com/gorilla/mux"

// SetupRoutes ...
func SetupRoutes(router *mux.Router, controller *TagController) {
	router.HandleFunc("/todo_items/{item_id}/tags", controller.GetAll).Methods("GET")
	router.HandleFunc("/todo_items/{item_id}/tags", controller.Post).Methods("POST")
	router.HandleFunc("/todo_items/{item_id}/tags/{tag_id}", controller.Remove).Methods("DELETE")
	router.HandleFunc("/tags/{id}", controller.GetSingle).Methods("GET")
	router.HandleFunc("/tags/{id}", controller.Put).Methods("PUT")
	router.HandleFunc("/tags/{id}", controller.Delete).Methods("DELETE")
}

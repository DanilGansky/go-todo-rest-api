package http

import "github.com/gorilla/mux"

// SetupTodoItemRoutes ...
func SetupTodoItemRoutes(router *mux.Router, controller *TodoItemController) {
	router.HandleFunc("/todo_lists/{list_id}/todo_items", controller.GetAll).Methods("GET")
	router.HandleFunc("/todo_lists/{list_id}/todo_items", controller.Post).Methods("POST")
	router.HandleFunc("/todo_items/{id}", controller.GetSingle).Methods("GET")
	router.HandleFunc("/todo_items/{id}", controller.Put).Methods("PUT")
	router.HandleFunc("/todo_items/{id}", controller.Delete).Methods("DELETE")
}

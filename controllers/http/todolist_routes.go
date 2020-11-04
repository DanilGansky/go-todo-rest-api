package http

import "github.com/gorilla/mux"

// SetupTodoListRoutes ...
func SetupTodoListRoutes(router *mux.Router, controller *TodoListController) {
	router.HandleFunc("/users/{user_id}/todo_lists", controller.GetAll).Methods("GET")
	router.HandleFunc("/users/{user_id}/todo_lists", controller.Post).Methods("POST")
	router.HandleFunc("/todo_lists/{id}", controller.GetSingle).Methods("GET")
	router.HandleFunc("/todo_lists/{id}", controller.Put).Methods("PUT")
	router.HandleFunc("/todo_lists/{id}", controller.Delete).Methods("DELETE")
}

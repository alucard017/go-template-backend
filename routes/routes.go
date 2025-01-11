package routes

import (
	"github.com/alucard777/go-todo-backend/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/tasks", controller.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", controller.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/tasks/{id}", controller.UpdateTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", controller.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", controller.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteTasks/", controller.DeleteTasks).Methods("DELETE", "OPTIONS")

	return router
}

package router

import (
	"github.com/Ash7540/todo-golang/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/task", middleware.GetAllToDoList).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tasks", middleware.CreateToDoList).Methods("POST", "OPTIONS")
	// router.HandleFunc("/api/task/{id}", middleware.GetToDoList).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/task/{id}", middleware.UpdateToDoList).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteToDoList).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTasks", middleware.DeleteAllList).Methods("DELETE", "OPTIONS")

	return router
}

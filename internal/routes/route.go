package routes

import (
    "github.com/boysimon10/go-mysql-crud/internal/handlers"
    "github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router, handler *handlers.BookHandler) {
    router.HandleFunc("/books", handler.CreateBook).Methods("POST")
    router.HandleFunc("/books", handler.GetAllBooks).Methods("GET")
    router.HandleFunc("/books/{id}", handler.GetBook).Methods("GET")
    router.HandleFunc("/books/{id}", handler.UpdateBook).Methods("PUT")
    router.HandleFunc("/books/{id}", handler.DeleteBook).Methods("DELETE")
}
package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/boysimon10/go-mysql-crud/config"
    "github.com/boysimon10/go-mysql-crud/internal/database"
    "github.com/boysimon10/go-mysql-crud/internal/handlers"
    "github.com/boysimon10/go-mysql-crud/internal/repository"
    "github.com/boysimon10/go-mysql-crud/internal/routes"
    "github.com/gorilla/mux"
)

func main() {
    // Load config
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal(err)
    }

    // Initialize database
    db, err := database.InitDB(cfg.GetDSN())
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Initialize repository
    bookRepo := repository.NewBookRepository(db)

    // Initialize handler
    bookHandler := handlers.NewBookHandler(bookRepo)

    // Setup router
    router := mux.NewRouter()
    routes.SetupRoutes(router, bookHandler)

    // Start server
    serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
    log.Printf("Server starting on port %s", cfg.ServerPort)
    log.Fatal(http.ListenAndServe(serverAddr, router))
}
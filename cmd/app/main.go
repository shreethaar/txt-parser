package main

import (
    "log"
    "net/http"
    "github.com/go-chi/chi/v5"
    "txtparser/internal/parser"
    "txtparser/internal/templates"
)

func main() {
    r := chi.NewRouter()
    SetupRoutes(r) 
    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

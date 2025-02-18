package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "txt-parser/internal/parser"
    "txt-parser/internal/templates"
)

func SetupRoutes(r *chi.Mux) {
    r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        content, err := parser.ReadFile("assets/data.txt")
        if err != nil {
            http.Error(w, "Failed to read file", http.StatusInternalServerError)
            return
        }
        templates.RenderHome(w, content)
    })


    r.Get("/edit", func(w http.ResponseWriter, r *http.Request) {
        content, err := parser.ReadFile("assets/data.txt")
        if err != nil {
            http.Error(w, "Failed to read file", http.StatusInternalServerError)
            return
        }
        templates.RenderEdit(w, content)
    })
    r.Post("/update", func(w http.ResponseWriter, r *http.Request) {
        content := r.FormValue("content")
        err := parser.UpdateFile("assets/data.txt", content)
        if err != nil {
            http.Error(w, "Failed to update file", http.StatusInternalServerError)
            return
        }
        http.Redirect(w, r, "/", http.StatusSeeOther)
    })
}

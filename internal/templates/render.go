package templates

import (
    "html/template"
    "net/http"
    "github.com/a-h/templ"
    "txtparser/views/home"
    "txtparser/views/edit"
    "txtparser/views/layouts"
)

func RenderHome(w http.ResponseWriter, content templ.Component) {
    templ.Handler(home.Home(content)).ServeHTTP(w, nil)
}

func RenderEdit(w http.ResponseWriter, content templ.Component) {
    templ.Handler(edit.Edit(content)).ServeHTTP(w, nil)
}

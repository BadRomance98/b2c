package controllers

import (
	"html/template"
	"net/http"
)

func Nopage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/404.html")
	t.Execute(w, nil)
}

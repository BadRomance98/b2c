package controllers

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/home.html", "views/header.tpl")
	pageNotFound(w, err)
	t.Execute(w, nil)
}

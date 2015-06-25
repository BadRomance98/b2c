package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/home.html", "views/header.tpl")
	if err != nil {
		log.Fatalf("找不到页面,错误代码:%s", err)
		http.Redirect(w, r, "/nopage", http.StatusFound)
	}
	t.Execute(w, nil)
}

package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func loginCookie(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("Uname")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func pageNotFound(w http.ResponseWriter, err error) {
	if err != nil {
		log.Println("找不到页面,错误代码:" + err.Error())
		t, _ := template.ParseFiles("views/404.html")
		t.Execute(w, nil)
	}
}

func CheckError(s string, err error) {
	if err != nil {
		log.Println(s + err.Error())
	}
}

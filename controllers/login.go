package controllers

import (
	"go-YTP/models"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {
	tel, err := template.ParseFiles("views/login.html", "views/header.tpl")
	pageNotFound(w, err)
	isExit := strings.EqualFold(r.FormValue("exit"), "true")
	if isExit {
		cookie, _ := r.Cookie("Uname")
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	}

	_, err = r.Cookie("Uname")
	if err == nil {
		http.Redirect(w, r, "/manage", http.StatusFound)
	}

	data := make(map[string]string)
	tel.Execute(w, data)

}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	login, err := template.ParseFiles("views/login.html", "views/header.tpl")
	pageNotFound(w, err)

	uname := r.PostFormValue("uname")
	pwd := r.PostFormValue("pwd")
	checkbox := r.PostFormValue("autoLogin") == "on"

	user, has, err := models.GetUserInfo(uname)
	if err != nil {
		log.Println("查询用户失败:" + err.Error())
	}

	if has {
		if strings.EqualFold(user.Passwd, pwd) {
			if checkbox {
				newCookie := http.Cookie{
					Name:   http.CanonicalHeaderKey("uname"),
					Value:  uname,
					MaxAge: 1<<31 - 1,
				}
				http.SetCookie(w, &newCookie)
			} else {
				newCookie := http.Cookie{
					Name:   http.CanonicalHeaderKey("uname"),
					Value:  uname,
					MaxAge: 0,
				}
				http.SetCookie(w, &newCookie)
			}
			http.Redirect(w, r, "/manage", http.StatusFound)
		} else {
			data["error"] = "密码错误"
			data["uname"] = uname
			login.Execute(w, data)
		}
	} else {
		data["error"] = "用户不存在"
		login.Execute(w, data)
	}
}

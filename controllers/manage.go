package controllers

import (
	"go-YTP/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Manage(w http.ResponseWriter, r *http.Request) {
	loginCookie(w, r)
	data := make(map[string]interface{})
	manage, err := template.ParseFiles("views/manage.html", "views/manage.tpl", "views/header.tpl")
	pageNotFound(w, err)
	p, err := strconv.Atoi(r.FormValue("p"))
	if err != nil {
		p = 0
	}
	data["pagesize"], err = models.GetNewsCount()
	if err != nil {
		log.Println("数据库连接失败:" + err.Error())
	}

	data["newslimit"], err = models.GetNewsLimit(p)
	if err != nil {
		log.Println("数据库连接失败:" + err.Error())
	}

	data["isNews"] = true
	data["isMedia"] = false
	manage.Execute(w, data)

}

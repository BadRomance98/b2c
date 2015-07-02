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
	var page = [10]int{1}
	manage, err := template.ParseFiles("views/manage.html", "views/manage.tpl", "views/header.tpl")
	pageNotFound(w, err)
	p, err := strconv.Atoi(r.FormValue("p"))
	if err != nil {
		p = 0
	} else {
		p = p - 1
	}
	pagesize, err := models.GetNewsCount()
	if err != nil {
		log.Println("数据库连接失败:" + err.Error())
	}

	s := strconv.FormatInt(pagesize, 10)
	intpage, _ := strconv.Atoi(s)
	for i := 0; i < (((intpage - 1) / 10) + 1); i++ {
		page[i] = i + 1
	}

	newsList, err := models.GetNewsLimit(p)
	if err != nil {
		log.Println("数据库连接失败:" + err.Error())
	}
	for li, newsone := range newsList {
		if len(newsone.Content) > 20 {
			newsList[li].Content = string([]rune(newsone.Content)[0:17]) + "..."
		}
	}

	data["pagenum"] = p + 1
	data["pagesize"] = page
	data["newslimit"] = newsList
	data["isNews"] = true
	data["isMedia"] = false
	manage.Execute(w, data)

}

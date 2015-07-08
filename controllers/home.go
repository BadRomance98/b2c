package controllers

import (
	"encoding/json"
	"go-YTP/models"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/home.html", "views/info_header.tpl", "views/foot.tpl")
	pageNotFound(w, err)
	t.Execute(w, nil)
}

func HomeAjaxImg(w http.ResponseWriter, r *http.Request) {
	mediaList, err := models.GetMediaHomeList()
	CheckError("首页图片获取,数据库连接失败", err)

	newsList, err := models.GetNewsHomeList()
	CheckError("首页图片获取,数据库连接失败", err)

	rd := models.ReturnData{Media: mediaList, News: newsList}
	js, err := json.Marshal(rd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func HomeTerms(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/terms.html", "views/info_header.tpl", "views/foot.tpl", "views/header.tpl")
	pageNotFound(w, err)
	data := make(map[string]interface{})
	data["isTerm"] = true
	t.Execute(w, data)
}
func HomeJoin(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/join.html", "views/info_header.tpl", "views/foot.tpl", "views/header.tpl")
	pageNotFound(w, err)
	data := make(map[string]interface{})
	data["isJoin"] = true
	t.Execute(w, data)
}
func HomeContact(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/contact.html", "views/info_header.tpl", "views/foot.tpl", "views/header.tpl")
	pageNotFound(w, err)
	data := make(map[string]interface{})
	data["isContact"] = true
	t.Execute(w, data)
}
func HomePolicy(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/policy.html", "views/info_header.tpl", "views/foot.tpl", "views/header.tpl")
	pageNotFound(w, err)
	data := make(map[string]interface{})
	data["isPolicy"] = true
	t.Execute(w, data)
}
func Homemedia(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/media.html", "views/info_header.tpl", "views/foot.tpl", "views/header.tpl")
	pageNotFound(w, err)
	data := make(map[string]interface{})
	data["isContact"] = true
	t.Execute(w, data)
}

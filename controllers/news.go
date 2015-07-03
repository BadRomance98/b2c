package controllers

import (
	"go-YTP/models"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func NewsAdd(w http.ResponseWriter, r *http.Request) {
	loginCookie(w, r)
	newsAdd, err := template.ParseFiles("views/news_add.html", "views/manage.tpl", "views/header.tpl")
	pageNotFound(w, err)

	data := make(map[string]interface{})
	data["error"] = ""
	id := r.FormValue("id")
	data["id"] = id
	if !strings.EqualFold(id, "") {
		intId, _ := strconv.ParseInt(id, 10, 64)
		news, err := models.SelectNews(intId)
		if err != nil {
			log.Println("数据库查询失败:" + err.Error())
			http.Redirect(w, r, "/manage", http.StatusFound)
			return
		}
		data["news"] = news
	}
	data["isNews"] = true
	data["isMedia"] = false
	newsAdd.Execute(w, data)
}

func NewsAddPost(w http.ResponseWriter, r *http.Request) {
	loginCookie(w, r)
	newsAdd, err := template.ParseFiles("views/manage.html", "views/manage.tpl", "views/header.tpl")
	pageNotFound(w, err)

	id := r.FormValue("id")
	if !strings.EqualFold(id, "") {
		intId, _ := strconv.ParseInt(id, 10, 64)
		news := new(models.News)
		news.OrderBy, _ = strconv.Atoi(r.FormValue("orderby"))
		news.Status, _ = strconv.Atoi(r.FormValue("status"))
		err = models.UpdateNews(intId, news)
		if err != nil {
			log.Println("新闻更新失败")
		}
		http.Redirect(w, r, "/manage", http.StatusFound)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	hrefurl := r.FormValue("hrefurl")
	if !strings.Contains(hrefurl, "http") {
		hrefurl = "http://" + hrefurl
	}

	file, fileHead, err := r.FormFile("picture")

	if err != nil {
		log.Println("文件读取失败：" + err.Error())
		news := models.News{Title: title, Content: content, HrefUrl: hrefurl}
		data := make(map[string]interface{})
		data["news"] = news
		data["error"] = "请上传图片"
		newsAdd.Execute(w, data)
	}

	fileName := strconv.FormatInt(time.Now().Unix(), 10) + "_" + fileHead.Filename
	f, err := os.OpenFile("upload/image/news/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	io.Copy(f, file)
	if err != nil {
		log.Println("文件夹读取失败" + err.Error())
	}
	defer file.Close()
	defer f.Close()

	news := new(models.News)
	news.Title = title
	news.Content = content
	news.HrefUrl = hrefurl
	news.PictureUrl = fileName
	news.SubDate = time.Now()
	news.Status = 1
	news.OrderBy = 0
	err = models.InsertNews(news)
	if err != nil {
		log.Println("新闻添加失败:" + err.Error())
	}
	http.Redirect(w, r, "/manage", http.StatusFound)
	return
}

func NewsDel(w http.ResponseWriter, r *http.Request) {
	loginCookie(w, r)

	id := r.FormValue("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println("数据转换失败:" + err.Error())
	}

	err = models.DeleteNews(intId)
	if err != nil {
		log.Println("新闻删除失败:" + err.Error())
	}

	http.Redirect(w, r, "/manage", http.StatusFound)
	return
}

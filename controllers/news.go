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
	newsAdd.Execute(w, data)
}

func NewsAddPost(w http.ResponseWriter, r *http.Request) {
	loginCookie(w, r)
	newsAdd, err := template.ParseFiles("views/manage.html", "views/manage.tpl", "views/header.tpl")
	pageNotFound(w, err)

	id := r.FormValue("id")
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
	fileName := strconv.FormatInt(int64(time.Now().Second()), 10) + "_" + fileHead.Filename
	f, err := os.OpenFile("upload/image/news/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	io.Copy(f, file)
	if err != nil {
		log.Println("文件夹读取失败" + err.Error())
	}
	defer file.Close()
	defer f.Close()

	err = models.SaveOrUpdateNews(id, title, content, hrefurl, fileName)
	if err != nil {
		log.Println("新闻数据处理失败:" + err.Error())
	}
	http.Redirect(w, r, "/manage", http.StatusFound)
}

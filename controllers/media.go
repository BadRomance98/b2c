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

func MediaAdd(w http.ResponseWriter, r *http.Request) {
	loginCookie(w, r)
	mediaAdd, err := template.ParseFiles("views/media_add.html", "views/manage.tpl", "views/header.tpl")
	pageNotFound(w, err)

	data := make(map[string]interface{})
	data["error"] = ""
	id := r.FormValue("id")
	data["id"] = id
	if !strings.EqualFold(id, "") {
		intId, _ := strconv.ParseInt(id, 10, 64)
		media, err := models.SelectMedia(intId)
		if err != nil {
			log.Println("数据库查询失败:" + err.Error())
			http.Redirect(w, r, "/media", http.StatusFound)
			return
		}
		data["media"] = media
	}
	data["isNews"] = false
	data["isMedia"] = true
	mediaAdd.Execute(w, data)
}

func MediaAddPost(w http.ResponseWriter, r *http.Request) {
	loginCookie(w, r)
	mediaAdd, err := template.ParseFiles("views/manage_media.html", "views/manage.tpl", "views/header.tpl")
	pageNotFound(w, err)

	id := r.FormValue("id")
	if !strings.EqualFold(id, "") {
		intId, _ := strconv.ParseInt(id, 10, 64)
		media := new(models.Media)
		media.OrderBy, _ = strconv.Atoi(r.FormValue("orderby"))
		media.Status, _ = strconv.Atoi(r.FormValue("status"))
		err = models.UpdateMedia(intId, media)
		if err != nil {
			log.Println("媒体更新失败")
		}
		http.Redirect(w, r, "/media", http.StatusFound)
		return
	}

	hrefurl := r.FormValue("hrefurl")
	if !strings.Contains(hrefurl, "http") {
		hrefurl = "http://" + hrefurl
	}

	file, fileHead, err := r.FormFile("picture")

	if err != nil {
		log.Println("文件读取失败：" + err.Error())
		media := models.Media{HrefUrl: hrefurl}
		data := make(map[string]interface{})
		data["media"] = media
		data["error"] = "请上传图片"
		mediaAdd.Execute(w, data)
	}

	fileName := strconv.FormatInt(time.Now().Unix(), 10) + "_" + fileHead.Filename
	f, err := os.OpenFile("upload/image/media/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	io.Copy(f, file)
	if err != nil {
		log.Println("文件夹读取失败" + err.Error())
	}
	defer file.Close()
	defer f.Close()

	media := new(models.Media)
	media.HrefUrl = hrefurl
	media.PictureUrl = fileName
	media.SubDate = time.Now()
	media.Status = 1
	media.OrderBy = 0
	err = models.InsertMedia(media)
	if err != nil {
		log.Println("媒体添加失败:" + err.Error())
	}
	http.Redirect(w, r, "/media", http.StatusFound)
	return
}

func MediaDel(w http.ResponseWriter, r *http.Request) {
	loginCookie(w, r)

	id := r.FormValue("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println("媒体转换失败:" + err.Error())
	}

	err = models.DeleteMedia(intId)
	if err != nil {
		log.Println("媒体删除失败:" + err.Error())
	}

	http.Redirect(w, r, "/media", http.StatusFound)
	return
}

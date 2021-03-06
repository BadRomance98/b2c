package main

import (
	// "html/template"
	"go-YTP/conf"
	"go-YTP/controllers"
	_ "go-YTP/models"
	"net/http"
	"runtime"
)

var HttpPort string

func init() {
	port, _ := conf.Cfg.GetValue("", "httpport")
	HttpPort = ":" + port
}

func main() {
	//多核运行
	runtime.GOMAXPROCS(runtime.NumCPU())
	//静态文件
	http.Handle("/css/", http.FileServer(http.Dir("statics")))
	http.Handle("/js/", http.FileServer(http.Dir("statics")))
	http.Handle("/img/", http.FileServer(http.Dir("statics")))
	http.Handle("/fonts/", http.FileServer(http.Dir("statics")))
	http.Handle("/image/", http.FileServer(http.Dir("upload")))

	//路由
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/loginPost", controllers.LoginPost)

	http.HandleFunc("/manage", controllers.Manage)

	http.HandleFunc("/news/add", controllers.NewsAdd)
	http.HandleFunc("/news/addPost", controllers.NewsAddPost)
	http.HandleFunc("/news/delete", controllers.NewsDel)
	http.HandleFunc("/news/edit", controllers.NewsAdd)

	http.HandleFunc("/media", controllers.ManageMedia)
	http.HandleFunc("/media/add", controllers.MediaAdd)
	http.HandleFunc("/media/addPost", controllers.MediaAddPost)
	http.HandleFunc("/media/delete", controllers.MediaDel)
	http.HandleFunc("/media/edit", controllers.MediaAdd)

	http.HandleFunc("/home/ajax", controllers.HomeAjaxImg)

	http.HandleFunc("/terms", controllers.HomeTerms)
	http.HandleFunc("/join", controllers.HomeJoin)
	http.HandleFunc("/contact", controllers.HomeContact)
	http.HandleFunc("/policy", controllers.HomePolicy)
	http.HandleFunc("/media_read", controllers.Homemedia)
	http.ListenAndServe(HttpPort, nil)
}

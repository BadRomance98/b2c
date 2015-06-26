package main

import (
	// "html/template"
	"go-YTP/conf"
	"go-YTP/controllers"
	// _ "go-YTP/models"
	"net/http"
)

var HttpPort string

func init() {
	port, _ := conf.Cfg.GetValue("", "httpport")
	HttpPort = ":" + port
}

func main() {
	//静态文件
	http.Handle("/css/", http.FileServer(http.Dir("statics")))
	http.Handle("/js/", http.FileServer(http.Dir("statics")))
	http.Handle("/img/", http.FileServer(http.Dir("statics")))
	http.Handle("/fonts/", http.FileServer(http.Dir("statics")))

	//路由
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/nopage", controllers.Nopage)

	http.ListenAndServe(HttpPort, nil)
}

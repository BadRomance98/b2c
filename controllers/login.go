package controllers

import (
	// "html/template"
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "just for test!")
}

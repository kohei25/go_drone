package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/kohei25/go_drone.git/config"
)

func viewIndexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("app/views/index.html")
	err := t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StartWebServer() error {
	http.HandleFunc("/", viewIndexHandler)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", config.Config.Address), nil)
}

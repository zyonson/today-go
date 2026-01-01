package controllers

import (
	"fmt"
	"net/http"
	"text/template"
	. "today-go/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "top", data)
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", ShowEventsAndPlaces)
	return http.ListenAndServe(":"+Config.Port, nil)
}

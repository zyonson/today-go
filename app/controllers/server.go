package controllers

import (
	"net/http"
	. "today-go/config"
)

func StartMainServer() error {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(Config.Static))))

	http.HandleFunc("/", ShowEventsAndPlaces)
	http.HandleFunc("/api/token", handleToken)
	http.HandleFunc("/top", handleTop)
	http.HandleFunc("/api/data", handleData)

	return http.ListenAndServe(":"+Config.Port, nil)
}

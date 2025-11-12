package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"today-go/config"
)

type EventList struct {
	Events []Event `json:"items"`
}

type Event struct {
	EventName   string `json:"summary"`
	DestiNation string `json:"location"`
}

type Place struct {
	Rating      float64 `json:"rating"`
	MapUri      string  `json:"googleMapsUri"`
	StoreSite   string  `json:"websiteUri"`
	DisplayName struct {
		Text string `json:"text"`
	} `json:"displayName"`
}

type PlacesResponse struct {
	Places []Place `json:"places"`
}

func index(w http.ResponseWriter, r *http.Request) {

	eventFile, err := os.Open("b.json")
	if err != nil {
		log.Fatal("ファイルを開けませんでした。", err)
	}
	defer eventFile.Close()

	eventBytes, err := io.ReadAll(eventFile)
	if err != nil {
		log.Fatal("ファイルの読み込みに失敗しました。", err)
	}

	var eventList EventList
	err = json.Unmarshal(eventBytes, &eventList)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("a.json")
	if err != nil {
		log.Fatal("ファイルを開けませんでした。", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("ファイルの読み込みに失敗しました。", err)
	}

	var places PlacesResponse
	err = json.Unmarshal(bytes, &places)
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]interface{}{
		"eventList": eventList,
		"places":    places,
		"apiKey":    config.Config.GoogleMapsAPIKey,
	}
	generateHTML(w, data, "top")
}

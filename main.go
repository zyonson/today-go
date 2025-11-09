package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

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

func main() {
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

	for _, p := range places.Places {
		fmt.Println("お店の名前:", p.DisplayName.Text, "お店の評価:", p.Rating, "お店の地図:", p.MapUri, "公式サイト:", p.StoreSite)
	}
}

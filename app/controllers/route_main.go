package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	. "today-go/config"
	. "today-go/models"
)

func index(w http.ResponseWriter, r *http.Request) {
	// イベント情報の読み込み
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

	// Google Places APIでお店情報を取得
	apiKey := Config.GoogleMapsAPIKey
	if apiKey == "" {
		log.Fatal("APIキーが設定されていません")
	}

	query := "レストラン " + eventList.Events[0].DestiNation
	endpoint := "https://maps.googleapis.com/maps/api/place/textsearch/json"

	params := url.Values{}
	params.Add("query", query)
	params.Add("key", apiKey)
	params.Add("language", "ja")

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	resp, err := http.Get(fullURL)
	if err != nil {
		log.Fatal("リクエスト失敗:", err)
	}
	defer resp.Body.Close()

	var result PlacesAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal("レスポンスの解析失敗:", err)
	}

	var places []Place
	for _, r := range result.Results {
		places = append(places, Place{
			Rating: r.Rating,
			MapUri: "https://www.google.com/maps/place/?q=place_id:" + r.PlaceID,
			DisplayName: DisplayName{
				Text: r.Name,
			},
		})
	}

	// テンプレートに渡すデータ
	data := map[string]interface{}{
		"eventList": eventList,
		"places":    PlacesResponse{Places: places},
		"apiKey":    apiKey,
	}

	generateHTML(w, data, "top")
}

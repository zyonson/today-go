package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	. "today-go/config"
	. "today-go/models"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func ShowEventsAndPlaces(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, Config.Static+"/templates/login.html")
}

var store = sessions.NewCookieStore([]byte(Config.SessionSecret))

func handleToken(w http.ResponseWriter, r *http.Request) {
	var req TokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	session, _ := store.Get(r, "session-name")
	session.Values["access_token"] = req.AccessToken
	session.Values["userId"] = req.Email

	if err := session.Save(r, w); err != nil {
		http.Error(w, "セッション保存に失敗しました", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func handleTop(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, Config.Static+"/templates/top.html")
}

func handleData(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	accessToken, _ := session.Values["access_token"].(string)
	token := &oauth2.Token{AccessToken: accessToken}
	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))
	srv, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		http.Error(w, "カレンダーサービスの初期化に失敗しました", http.StatusInternalServerError)
		return
	}

	now := time.Now()
	currentTime := now.Format(time.RFC3339)
	endOfDay := time.Date(
		now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location(),
	)
	endStr := endOfDay.Format(time.RFC3339)

	events, err := srv.Events.List("primary").ShowDeleted(false).SingleEvents(true).TimeMin(currentTime).TimeMax(endStr).MaxResults(10).OrderBy("startTime").Do()

	if err != nil {
		http.Error(w, "イベントの取得に失敗しました", http.StatusInternalServerError)
		return
	}

	if len(events.Items) == 0 {
		http.Error(w, "今日のイベントが見つかりませんでした。", http.StatusNotFound)
		return
	}

	apiKey := Config.GoogleMapsAPIKey
	if apiKey == "" {
		http.Error(w, "Google Maps APIキーが設定されていません", http.StatusInternalServerError)
		return
	}

	var eventList []Event
	for _, item := range events.Items {
		eventName := item.Summary
		destiNation := strings.Split(item.Location, ",")[0]

		eventList = append(eventList, Event{
			EventName:   eventName,
			DestiNation: destiNation,
		})
	}

	query := "レストラン " + eventList[0].DestiNation
	endpoint := "https://maps.googleapis.com/maps/api/place/textsearch/json"

	params := url.Values{}
	params.Add("query", query)
	params.Add("key", apiKey)
	params.Add("language", "ja")

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	resp, err := http.Get(fullURL)
	if err != nil {
		http.Error(w, "Google Maps APIリクエストに失敗しました", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var result PlacesAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		http.Error(w, "Google Maps APIレスポンスの解析に失敗しました", http.StatusInternalServerError)
		return
	}

	var places []Place
	for _, r := range result.Results {
		places = append(places, Place{
			Rating: r.Rating,
			MapUri: fmt.Sprintf("https://www.google.com/maps/search/?api=1&query_place_id=%s", r.PlaceID),
			DisplayName: DisplayName{
				Text: r.Name,
			},
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"eventList": eventList,
		"places":    PlacesResponse{Places: places},
		"apiKey":    apiKey,
	})
}

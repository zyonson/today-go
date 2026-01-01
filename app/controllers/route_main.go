package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"time"

	. "today-go/config"
	. "today-go/models"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func getClient(config *oauth2.Config) *http.Client {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("ブラウザで以下のURLを開いてください:\n%v\n", authURL)

	_ = openBrowser(authURL)

	codeCh := make(chan string)
	srv := &http.Server{Addr: ":8082"}

	http.HandleFunc("/cognite", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		fmt.Printf("認証コード検証前までの処理です。")
		if code != "" {
			fmt.Fprintln(w, "認証が完了しました！このウィンドウを閉じてください。")
			codeCh <- code
		} else {
			http.Error(w, "認証コードが見つかりませんでした", http.StatusBadRequest)
		}
	})

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("サーバー起動エラー: %v", err)
		}
	}()

	code := <-codeCh

	ctx, cancel := context.WithTimeout(context.Background(), 1000)
	defer cancel()
	_ = srv.Shutdown(ctx)

	tok, err := config.Exchange(context.TODO(), code)
	if err != nil {
		log.Fatalf("トークン取得失敗: %v", err)
	}
	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("トークンを保存しました: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("トークンの保存に失敗しました: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func openBrowser(url string) error {
	var cmd string
	var args []string

	switch {
	case "windows" == os.Getenv("OS"):
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin" == os.Getenv("GOOS"):
		cmd = "open"
		args = []string{url}
	default:
		cmd = "xdg-open"
		args = []string{url}
	}
	return exec.Command(cmd, args...).Start()
}

func ShowEventsAndPlaces(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		http.Error(w, "クレデンシャルファイルの読み込みに失敗しました", http.StatusInternalServerError)
		return
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		http.Error(w, "OAuth設定の解析に失敗しました", http.StatusInternalServerError)
		return
	}
	client := getClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		http.Error(w, "カレンダーサービスの初期化に失敗しました", http.StatusInternalServerError)
		return
	}

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	//fmt.Printf("json結果 %s", events)
	if err != nil {
		http.Error(w, "イベントの取得に失敗しました", http.StatusInternalServerError)
		return
	}

	if len(events.Items) == 0 {
		http.Error(w, "今後のイベントが見つかりませんでした", http.StatusNotFound)
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
		destiNation := item.Location

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

	data := map[string]interface{}{
		"eventList": eventList,
		"places":    PlacesResponse{Places: places},
		"apiKey":    apiKey,
	}

	generateHTML(w, data, "top")
}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigList struct {
	Port             string
	Static           string
	GoogleMapsAPIKey string
	SessionSecret    string
}

var Config ConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("環境変数ファイル(.env)の読み込みに失敗しました:", err)
	}

	Config = ConfigList{
		Port:             os.Getenv("PORT"),
		Static:           os.Getenv("STATIC"),
		GoogleMapsAPIKey: os.Getenv("GOOGLE_MAPS_API_KEY"),
		SessionSecret:    os.Getenv("SESSION_SECRET"),
	}
}

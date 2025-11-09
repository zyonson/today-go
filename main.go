package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Place struct {
	DisplayName struct {
		Text string `json:"text"`
	} `json:"displayName"`
}

type PlacesResponse struct {
	Places []Place `json:"places"`
}

func main() {
	data := `{
      "places": [
        {
          "name": "places/ChIJL67AxKiMGGARXNGDs4npDHU",
          "id": "ChIJL67AxKiMGGARXNGDs4npDHU",
          "types": [
            "chinese_restaurant",
            "restaurant",
            "food",
            "point_of_interest",
            "establishment"
          ],
          "utcOffsetMinutes": 540,
          "adrFormatAddress": "\u003cspan class=\"country-name\"\u003e日本\u003c/span\u003e、\u003cspan class=\"postal-code\"\u003e〒150-0041\u003c/span\u003e \u003cspan class=\"region\"\u003e東京都\u003c/span\u003e\u003cspan class=\"street-address\"\u003e渋谷区神南１丁目１６−３ ブル・ヴァールビル B1\u003c/span\u003e",
          "businessStatus": "OPERATIONAL",
          "priceLevel": "PRICE_LEVEL_INEXPENSIVE",
          "userRatingCount": 1344,
          "iconMaskBaseUri": "https://maps.gstatic.com/mapfiles/place_api/icons/v2/restaurant_pinlet",
          "iconBackgroundColor": "#FF9E67",
          "displayName": {
            "text": "本格中華料理 陳家私菜 渋谷店",
            "languageCode": "ja"
          },
          "primaryTypeDisplayName": {
            "text": "中華料理店",
            "languageCode": "ja"
          },
          "takeout": true,
          "delivery": true,
          "dineIn": true,
          "curbsidePickup": true,
          "reservable": true,
          "servesBreakfast": false,
          "servesLunch": true,
          "servesDinner": true,
          "servesBeer": true,
          "servesWine": true,
          "servesBrunch": true,
          "servesVegetarianFood": false,
          "primaryType": "chinese_restaurant",
          "shortFormattedAddress": "渋谷区神南１丁目１６−３ ブル・ヴァールビル B1",
          "editorialSummary": {
            "text": "「元祖頂天石焼麻婆豆腐」と「頂天石焼麻婆刀削麺」が看板メニューの中華料理店。黄色い壁と大きな窓に囲まれた明るいダイニングスペース。",
            "languageCode": "ja"
          }
        },
        {
          "name": "places/ChIJe7uhwamMGGAR3r-Wrg2--fQ",
          "id": "ChIJe7uhwamMGGAR3r-Wrg2--fQ",
          "types": [
            "asian_restaurant",
            "chinese_restaurant",
            "restaurant",
            "food",
            "point_of_interest",
            "establishment"
          ],
          "nationalPhoneNumber": "03-3461-4220",
          "internationalPhoneNumber": "+81 3-3461-4220",
          "formattedAddress": "日本、〒150-0043 東京都渋谷区道玄坂２丁目２５−１８",
    
          "plusCode": {
            "globalCode": "8Q7XMM5X+Q3",
            "compoundCode": "MM5X+Q3 日本、東京都渋谷区"
          },
          "location": {
            "latitude": 35.659469,
            "longitude": 139.69768100000002
          },
          "viewport": {
            "low": {
              "latitude": 35.6581447197085,
              "longitude": 139.69636671970849
            },
            "high": {
              "latitude": 35.660842680291495,
              "longitude": 139.69906468029149
            }
          },
          "rating": 4,
          "googleMapsUri": "https://maps.google.com/?cid=17652349180428337118&g_mp=Cidnb29nbGUubWFwcy5wbGFjZXMudjEuUGxhY2VzLlNlYXJjaFRleHQQAhgBIAA",
          "websiteUri": "https://www.instagram.com/reikyo_dougenzaka/",
          "utcOffsetMinutes": 540,
          "adrFormatAddress": "\u003cspan class=\"country-name\"\u003e日本\u003c/span\u003e、\u003cspan class=\"postal-code\"\u003e〒150-0043\u003c/span\u003e \u003cspan class=\"region\"\u003e東京都\u003c/span\u003e\u003cspan class=\"street-address\"\u003e渋谷区道玄坂２丁目２５−１８\u003c/span\u003e",
          "businessStatus": "OPERATIONAL",
          "priceLevel": "PRICE_LEVEL_MODERATE",
          "userRatingCount": 1858,
          "iconMaskBaseUri": "https://maps.gstatic.com/mapfiles/place_api/icons/v2/restaurant_pinlet",
          "iconBackgroundColor": "#FF9E67",
          "displayName": {
            "text": "麗郷",
            "languageCode": "ja"
          },
          "primaryTypeDisplayName": {
            "text": "レストラン",
            "languageCode": "ja"
          },
          "takeout": true,
          "delivery": false,
          "dineIn": true,
          "reservable": true,
          "servesBreakfast": false,
          "servesLunch": true,
          "servesDinner": true,
          "servesBeer": true,
          "servesWine": false,
          "servesVegetarianFood": false,
          "primaryType": "restaurant",
          "shortFormattedAddress": "渋谷区道玄坂２丁目２５−１８",
          "editorialSummary": {
            "text": "魚介料理、餃子、こってりとした肉料理などの台湾料理を楽しめる、居心地の良いレンガ造りのレストラン。",
            "languageCode": "ja"
          }
        }
      ]
  }`

	var places PlacesResponse
	err := json.Unmarshal([]byte(data), &places)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range places.Places {
		fmt.Println("お店の名前:", p.DisplayName.Text)
	}
}

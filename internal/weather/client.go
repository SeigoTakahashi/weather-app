package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// レスポンスのJSONデータをWeatherResponse構造体に変換する関数
func ParseWeatherResponse(data []byte) (WeatherResponse, error) {
	var weatherResponse WeatherResponse
	err := json.Unmarshal(data, &weatherResponse)
	if err != nil {
		return WeatherResponse{}, fmt.Errorf("レスポンスのパースに失敗しました: %v", err)
	}
	return weatherResponse, nil
}

// 指定された都市コードとAPIキーを使用してOpenWeatherMap APIにGETリクエストを送信し、
// レスポンスをWeatherResponse構造体として返す。
func GetRequestToWeatherAPI(cityCode string, apiKey string) WeatherResponse {
	q := url.QueryEscape(cityCode)
	appid := url.QueryEscape(apiKey)
	
	targetUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", q, appid)

	// リクエストオブジェクトを作成 ("GET" を指定)
	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		fmt.Printf("リクエストの作成に失敗しました: %v\n", err)
		return WeatherResponse{}
	}

	// クライアントを使ってリクエストを送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("GETリクエストの送信中にエラーが発生しました: %v\n", err)
		return WeatherResponse{}
	}
	defer resp.Body.Close() // レスポンスボディを必ずクローズ

	// レスポンスのステータスコードを確認
	fmt.Printf("ステータスコード: %d\n", resp.StatusCode)

	// レスポンスボディを読み込む 
	body, err := io.ReadAll(resp.Body) 
	if err != nil {
		fmt.Printf("レスポンスボディの読み込み中にエラーが発生しました: %v\n", err)
		return WeatherResponse{}
	}

	// レスポンスを構造体にパース
	weatherResponse, err := ParseWeatherResponse(body)
	if err != nil {
		fmt.Printf("レスポンスのパースに失敗しました: %v\n", err)
		return WeatherResponse{}
	}
	
	return weatherResponse
}


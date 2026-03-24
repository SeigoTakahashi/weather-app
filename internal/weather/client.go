package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// レスポンスのJSONデータをWeatherResponse構造体に変換する関数
func ParseWeatherResponse(data []byte) WeatherResult {
	var weatherResponse WeatherResponse
	err := json.Unmarshal(data, &weatherResponse)
	if err != nil {
		return WeatherResult{Err: fmt.Errorf("レスポンスのパースに失敗しました: %v", err)}
	}
	return WeatherResult{Response: weatherResponse, Err: nil}
}

// 指定された都市コードとAPIキーを使用してOpenWeatherMap APIにGETリクエストを送信し、
// レスポンスをWeatherResponse構造体として返す。
func GetRequestToWeatherAPI(cityCode string, apiKey string) WeatherResult {
	q := url.QueryEscape(cityCode)
	appid := url.QueryEscape(apiKey)
	
	targetUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric&lang=ja", q, appid)

	// リクエストオブジェクトを作成 ("GET" を指定)
	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		fmt.Printf("リクエストの作成に失敗しました: %v\n", err)
		return WeatherResult{Err: fmt.Errorf("リクエストの作成に失敗しました: %v", err)}
	}

	// クライアントを使ってリクエストを送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("GETリクエストの送信中にエラーが発生しました: %v\n", err)
		return WeatherResult{Err: fmt.Errorf("GETリクエストの送信中にエラーが発生しました: %v", err)}
	}
	defer resp.Body.Close() // レスポンスボディを必ずクローズ

	// レスポンスボディを読み込む 
	body, err := io.ReadAll(resp.Body) 
	if err != nil {
		fmt.Printf("レスポンスボディの読み込み中にエラーが発生しました: %v\n", err)
		return WeatherResult{Err: fmt.Errorf("レスポンスボディの読み込み中にエラーが発生しました: %v", err)}
	}

	// レスポンスを構造体にパース
	weatherResult := ParseWeatherResponse(body)
	if weatherResult.Err != nil {
		fmt.Printf("レスポンスのパースに失敗しました: %v\n", weatherResult.Err)
		return WeatherResult{Err: fmt.Errorf("レスポンスのパースに失敗しました: %v", weatherResult.Err)}
	}
	
	return weatherResult
}


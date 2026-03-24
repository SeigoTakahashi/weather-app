package main

import (
	"os"
	"weather-app/infra"
	"weather-app/internal/weather"
)

func main() {
	// 環境変数の初期化
	infra.Initialize()

	// リクエストテスト
	weatherResponse := weather.GetRequestToWeatherAPI("Ichihara", os.Getenv("WEATHER_API_KEY"))
	// 結果の表示
	if weatherResponse.Name != "" {
		println("都市名:", weatherResponse.Name)
		println("天気:", weatherResponse.Weather[0].Icon)
		println("気温:", weatherResponse.Main.Temp)
	} else {
		println("天気情報の取得に失敗しました。")
	}

}
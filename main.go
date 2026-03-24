package main

import (
	"fmt"
	"os"
	"sync"
	"weather-app/infra"
	"weather-app/internal/service"
	"weather-app/internal/weather"
)

// 指定された都市の天気情報を取得し、チャネルに送信する関数
func fetchWeather(city string, c chan<- weather.WeatherResult, wg *sync.WaitGroup) {
	defer wg.Done() // 関数が終わったら「一人終わったよ」と報告

	// OpenWeatherMap API を叩く
	weatherResult := weather.GetRequestToWeatherAPI(city, os.Getenv("WEATHER_API_KEY"))
	if weatherResult.Err != nil {
		fmt.Printf("エラー (%s): %v\n", city, weatherResult.Err)
		c <- weatherResult
		return
	}

	// 結果をチャネルに送信
	c <- weatherResult
}

func main() {
	// 環境変数の初期化
	infra.Initialize()

	// 都市のリスト
	cities := []string{"Chiba,JP", "Ichihara,JP", "Tokyo,JP"}
	results := make(chan weather.WeatherResult, len(cities)) // 都市の数だけバッファを確保

	// WaitGroup を使って全てのゴルーチンの終了を待つ
	var wg sync.WaitGroup

	// 各都市の天気情報を非同期に取得
	for _, city := range cities {
		wg.Add(1) // これから一つずつ作業を始めることを報告
		go fetchWeather(city, results, &wg)
	}

	// すべての Goroutine が終わるのを待ってからチャネルを閉じる
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// チャネルが閉じるまで結果を受け取り続ける
	for res := range results {
		if res.Err != nil {
			fmt.Printf("エラー: %v\n", res.Err)
			continue
		}

		// 天気情報の表示
		fmt.Printf("----- %sの天気情報 -----\n", res.Response.Name)
		fmt.Printf("天候: %s\n", service.GetWeatherMessage(res.Response.Weather[0].Id))
		fmt.Printf("気温：%.1f度\n", res.Response.Main.Temp)
		fmt.Printf("体感気温：%.1f度\n", res.Response.Main.FeelsLike)
		fmt.Printf("湿度：%d%%\n", res.Response.Main.Humidity)
		fmt.Printf("風速：%.1fm/s\n", res.Response.Wind.Speed)
		advices := service.GetPracticalAdvice(&res.Response)
		fmt.Println("生活に役立つアドバイス:")
		for _, advice := range advices {
			fmt.Printf("- %s\n", advice)
		}
	}

}
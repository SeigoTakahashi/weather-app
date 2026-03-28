package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"weather-app/infra"
	"weather-app/internal/notifier"
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

		

		// 天気情報を一つの文字列変数にまとめる
		var b strings.Builder

		// 基本情報の書き込み
		fmt.Fprintf(&b, "----- *%sの天気情報* -----\n", res.Response.Name)
		fmt.Fprintf(&b, "*天候*: %s\n", service.GetWeatherMessage(res.Response.Weather[0].Id))
		fmt.Fprintf(&b, "*気温*: %.1f度\n", res.Response.Main.Temp)
		fmt.Fprintf(&b, "*体感気温*: %.1f度\n", res.Response.Main.FeelsLike)
		fmt.Fprintf(&b, "*湿度*: %d%%\n", res.Response.Main.Humidity)
		fmt.Fprintf(&b, "*風速*: %.1fm/s\n", res.Response.Wind.Speed)

		// アドバイス部分
		advices := service.GetPracticalAdvice(&res.Response)
		b.WriteString("*生活に役立つアドバイス*:\n")
		for _, advice := range advices {
			fmt.Fprintf(&b, "- %s\n", advice)
		}

		// 最終的な文字列を取得
		message := b.String()

		// 画像URLの生成
		iconCode := res.Response.Weather[0].Icon
		iconUrl := fmt.Sprintf("https://openweathermap.org/img/wn/%s@2x.png", iconCode)

		// Slackに送信
		err := notifier.PushMessageToSlack(os.Getenv("SLACK_BOT_USER_TOKEN"), "C0AL88UBYF7", message, iconUrl)
		if err != nil {
			fmt.Printf("Slackへの送信に失敗: %v\n", err)
		} else {
			fmt.Printf("%sの天気情報をSlackに送信しました。\n", res.Response.Name)
		}

	}

}
package service

import (
	"testing"
	"weather-app/internal/weather"
)

// 天候に応じたメッセージ作成のロジックテスト
func TestGetWeatherMessage(t *testing.T) {
	// テストケースの定義
	tests := []struct {
		name string // テストケースの名前
		id   int    // 天気ID
		want string // 期待されるメッセージ
	}{
		{
			name: "clear weather",
			id:   800,
			want: "快晴",
		},
		{
			name: "few clouds",
			id:   801,
			want: "雲が少ない",
		},
		{
			name: "atmosphere",
			id:   701,
			want: "ミスト",
		},
		{
			name: "unknown weather",
			id:   999,
			want: "不明な天気",
		},
	}

	// テストケースの実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWeatherMessage(tt.id); got != tt.want {
				t.Errorf("GetWeatherMessage(%d) = %v, want %v", tt.id, got, tt.want)
			}
		})
	}
}

// 天候に応じた実用的なアドバイスの生成ロジックテスト
func TestGetPracticalAdvice(t *testing.T) {
	// テストケースの定義
	tests := []struct {
		name string // テストケースの名前
		w    *weather.WeatherResponse // 入力となる天気データ
		want []string // 期待されるアドバイスのリスト
	}{
		{
			name: "hot and humid",
			w: &weather.WeatherResponse{
				Main: struct {
					Temp     float64 `json:"temp"`
					FeelsLike float64 `json:"feels_like"`
					Humidity int     `json:"humidity"`
				}{
					Temp: 30.0,
					FeelsLike: 35.0,
					Humidity: 80,
				},
				Wind: struct {
					Speed float64 `json:"speed"`
				}{
					Speed: 5.0,
				},
			},
			want: []string{
				"真夏日です。こまめな水分補給を忘れずに。",
				"蒸し暑さを感じそうです。除湿などで調整しましょう。",
				"やや風があります。体感温度は表示より低く感じられそうです。",
			},
		},
		{
			name: "cold and dry",
			w: &weather.WeatherResponse{
				Main: struct {
					Temp     float64 `json:"temp"`
					FeelsLike float64 `json:"feels_like"`
					Humidity int     `json:"humidity"`
				}{
					Temp: 0.0,
					FeelsLike: -5.0,
					Humidity: 30,
				},
				Wind: struct {
					Speed float64 `json:"speed"`
				}{
					Speed: 0.0,
				},
			},
			want: []string{
				"厳しい寒さです。防寒対策を万全に。路面凍結にも注意。",
				"乾燥して風邪を引きやすい環境です。加湿器を使いましょう。",
			},
		},
	}

	// テストケースの実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPracticalAdvice(tt.w); !equalStringSlices(got, tt.want) {
				t.Errorf("GetPracticalAdvice() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 文字列スライスの等価性をチェックするヘルパー関数
func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
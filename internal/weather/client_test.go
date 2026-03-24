package weather

import (
	"testing"
)

// 天気レスポンスのパース（解析）テスト
func TestUnmarshalWeather(t *testing.T) {
	// テストケースの定義
	tests := []struct {
		name string // テストケースの名前
		data []byte // JSONデータ
		wantError bool // エラーが期待されるかどうか
	}{
		{
			name: "valid weather response",
			data: []byte(`{"main":{"temp":20.5, "feels_like":19.8, "humidity":80}, "name":"Chiba", "weather":[{"id":800, "icon":"01d"}], "wind":{"speed":3.6}}`),
			wantError: false,
		},
		{
			name: "invalid weather response",
			data: []byte(`{"main":{"temp":"invalid", "feels_like":19.8, "humidity":80}, "name":"Chiba", "weather":[{"id":800, "icon":"01d"}], "wind":{"speed":3.6}}`),
			wantError: true,
		},
		{
			name: "minimal weather response",
			data: []byte(`{"main":{"temp":20.5}, "name":"Chiba"}`),
			wantError: false,
		},
	}
		
    // テストケースの実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			weatherResult := ParseWeatherResponse(tt.data)
			
			// エラーの有無が期待通りかチェック
            if (weatherResult.Err != nil) != tt.wantError {
                t.Errorf("Unmarshal() error = %v, wantErr %v", weatherResult.Err, tt.wantError)
            }
		})
	}   
}
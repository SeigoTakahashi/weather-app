package notifier

import (
	"fmt"
	"net/http"
	"strings"
)

func PushMessageToSlack(token string, channel string, message string, image_url string) error {
	url := "https://slack.com/api/chat.postMessage"

	data := `{"channel": "` + channel + `", "blocks": [{"type": "section", "text": {"type": "mrkdwn", "text": "` + message + `"}}, {"type": "image", "image_url": "` + image_url + `", "alt_text": "weather icon"}]}`
	bodyReader := strings.NewReader(data)

	// リクエストオブジェクトを作成 (まだ送信はされない)
	req, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		fmt.Printf("リクエストの作成に失敗しました: %v\n", err)
		return err
	}

	// ヘッダーを設定
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+token)

	// クライアントを使ってリクエストを送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("リクエストの送信中にエラーが発生しました: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	return nil
}
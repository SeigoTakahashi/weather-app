package infra

import (
	"log"

	"github.com/joho/godotenv"
)

// 初期化関数：
// - 環境変数の読み込み
func Initialize() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}
}
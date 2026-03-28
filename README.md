# 天気リマインダー

任意の地点の天気情報を毎日Slackに通知するアプリです。

---

## 🛠 技術スタック

- **Language**: Go 1.24.2
- **API Integration**:
  - Slack API
  - OpenWeatherMap API

- **Tooling**:
  - `godotenv`：環境変数管理

---

## 🔑 環境変数の設定

実行には以下の API キーが必要です。

### ローカルの場合

プロジェクトルートに `.env` ファイルを作成し、設定してください。

```env
# OpenWeatherMap API Key
WEATHER_API_KEY=your_openweathermap_api_key

# Slack API Key
SLACK_BOT_USER_TOKEN=your_slack_api_key
```

### Github Actions利用の場合

Github Secretに上記の変数を設定

---

## 🚀 開発環境の起動（Docker）

### 本番同等確認

GitHub Actionsの `schedule.yml` と同じ動きを確認したい場合：

```bash
# イメージをビルド
docker build -t weather-app .

# 実行（.envファイルがある場合）
docker run --rm --env-file .env weather-app
```

### 開発中にテストやLintを実行する（Compose利用）

`docker-compose.yml` を使って、開発環境（builderステージ）に入る場合：

```bash
# 1. ビルド
docker compose build

# 2. テストの実行
docker compose run --rm app go test -v ./...

# 3. Lintの実行
docker compose run --rm app golangci-lint run

# 4. 動作確認（コンテナ内で直接実行）
docker compose run --rm app go run main.go

# 5. コンテナの中に入って作業したい場合
docker compose run --rm app sh
```

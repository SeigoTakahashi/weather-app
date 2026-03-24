# --- Build Stage ---
FROM golang:1.24-alpine AS builder

# 必要なビルドツール（gitなど）をインストール
RUN apk add --no-cache git && \
    wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.64.5

WORKDIR /app

# 先に依存関係をコピーしてキャッシュを利用
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピーしてビルド
COPY . .
RUN go build -o /weather-app .

# --- Run Stage ---
FROM alpine:latest

# CA証明書のインストール（HTTPS通信を行うAPIを叩く場合に必須！）
RUN apk add --no-cache ca-certificates

WORKDIR /root/

# ビルドしたバイナリのみをコピー
COPY --from=builder /weather-app .

# 実行
CMD ["./weather-app"]
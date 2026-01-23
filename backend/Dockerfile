# Goの実行環境が入った軽量Linux(Alpine)をベースにする
FROM golang:1.25-alpine

# コンテナ内の作業ディレクトリを /app に設定
WORKDIR /app

# Airのインストール
RUN go install github.com/air-verse/air@latest

# 依存関係ファイルを先にコピー
COPY go.mod go.sum ./

# ライブラリをダウンロード
RUN go mod download

# 残りのソースコードをすべてコンテナにコピー
COPY . .

# Goアプリをビルドして "main" という実行ファイルを作る
RUN go build -o main .

# コンテナ起動時に、ビルドした "main" を実行する
CMD ["air", "-c", ".air.toml"]
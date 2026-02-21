FROM golang:1.24-alpine

# パッケージキャッシュを更新してからインストール
RUN apk update && apk add --no-cache git curl libc6-compat

# Air (Go 1.23と互換性のあるバージョンを指定)
RUN go install github.com/air-verse/air@v1.61.1

WORKDIR /app9ii98

# 依存関係のコピーとダウンロード
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# 先に Atlas をインストール（ソースコピーの影響を受けないように）
RUN curl -sSf https://atlasgo.sh | sh

# ソースコードをコピー
COPY backend/ .

# その他のツール
RUN go install github.com/aarondl/sqlboiler/v4@latest && \
    go install github.com/aarondl/sqlboiler/v4/drivers/sqlboiler-psql@latest
RUN go install github.com/google/wire/cmd/wire@latest

# Airが自動でビルドするため、ここでのビルドは不要
# RUN go build -o main .

CMD ["air", "-c", ".air.toml"]
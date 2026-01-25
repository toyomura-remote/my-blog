FROM golang:1.25-alpine

# git, curl に加えて libc6-compat をインストール
RUN apk add --no-cache git curl libc6-compat

# Air
RUN go install github.com/air-verse/air@latest

WORKDIR /app

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

RUN go build -o main .

CMD ["air", "-c", ".air.toml"]
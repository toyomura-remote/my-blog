# --- ステージ1: ビルド環境 ---
FROM golang:1.24.1-alpine AS builder

RUN apk update && apk add --no-cache make gcc g++ curl && \
    curl -sSf https://atlasgo.sh | sh

ENV GOPATH /go
ENV GO111MODULE=on
WORKDIR /go/src/main

# backend フォルダの中身をコピーするように修正
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# backend の全ソースをコピー
COPY backend/ .

# バイナリビルド (main.go は backend/直下にある)
RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -p $(nproc) -o /app/main -ldflags="-s -w" main.go

# --- ステージ2: 実行環境 ---
FROM alpine:3.19.1
RUN apk add --no-cache ca-certificates tzdata postgresql postgresql-client libpq gcompat
WORKDIR /app

COPY --from=builder /app/main /app/main
COPY --from=builder /usr/local/bin/atlas /usr/local/bin/atlas
COPY --from=builder /go/src/main/atlas.hcl /app/atlas.hcl
COPY --from=builder /go/src/main/db/migrations /app/db/migrations

EXPOSE 8080

CMD ["sh", "-c", "atlas migrate apply --url \"$DATABASE_URL\" --dir \"file://db/migrations\" && /app/main"]
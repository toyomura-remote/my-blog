# backend/Dockerfile
FROM golang:1.24.1-alpine AS builder
RUN apk update && apk add --no-cache gcc g++

WORKDIR /app

# すでに /backend の中にいるので、直下の go.mod を見る
COPY go.mod go.sum ./
RUN go mod download

# backend フォルダの中身（今いる場所）をすべてコピー
COPY . .

# cmd/main.go のパスを確認（backend/cmd/main.go なら cmd/main.go でOK）
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:3.19.1
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app

COPY --from=builder /app/main .
# マイグレーションファイル一式もコピー（一応持っておく）
COPY --from=builder /app/db/migrations ./db/migrations

EXPOSE 8080

CMD ["./main"]
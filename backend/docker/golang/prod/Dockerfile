# backend/Dockerfile
FROM golang:1.24.1-alpine AS builder
RUN apk update && apk add --no-cache gcc g++

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# アプリのビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

FROM alpine:3.19.1
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app

# バイナリだけをコピー（マイグレーションファイルも不要なら消してOK）
COPY --from=builder /app/main .

EXPOSE 8080

# アプリを起動するだけ！
CMD ["./main"]
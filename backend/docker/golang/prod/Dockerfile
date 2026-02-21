# backend/Dockerfile
FROM golang:1.24.1-alpine AS builder
RUN apk update && apk add --no-cache gcc g++

WORKDIR /app

# /backend のルートにいるので、直下の go.mod / go.sum をコピー
COPY go.mod go.sum ./
RUN go mod download

# backend フォルダの中身（今いる場所）をすべてコピー
COPY . .

# backend 直下に main.go があるので ./main.go を指定
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

FROM alpine:3.19.1
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app

# ビルドしたバイナリをコピー
COPY --from=builder /app/main .

EXPOSE 8080

# 実行
CMD ["./main"]
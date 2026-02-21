# backend/Dockerfile
FROM golang:1.24.1-alpine AS builder
RUN apk update && apk add --no-cache gcc g++

WORKDIR /app

# my-blog/ から見たパスでコピーする
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# backend の中身をすべて /app にコピー
COPY backend/ .

# backend 直下の main.go をビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

FROM alpine:3.19.1
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
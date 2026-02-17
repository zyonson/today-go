FROM golang:1.24.4 AS builder
WORKDIR /build
COPY . .
# ビルドディレクトリの外にバイナリを作成
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o today-go-binary .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=builder /build/today-go-binary ./today-go-binary
COPY .env .env
COPY app/views/ app/views/

RUN chmod +x ./today-go-binary
EXPOSE 80
CMD ["./today-go-binary"]
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o gin-server .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/gin-server .

EXPOSE 8080

CMD ["./gin-server"]

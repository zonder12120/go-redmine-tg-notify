# STAGE 1 BUILD

FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go test ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tg-bot ./cmd

# STAGE 2 RUN

FROM alpine:latest

RUN addgroup -S appgroup && adduser -S appuser -G appgroup  

WORKDIR /app

COPY --from=builder /app/tg-bot .

COPY --from=builder /app/cmd/.env .

RUN chown -R appuser:appgroup /app

USER appuser

CMD ["/app/tg-bot"]


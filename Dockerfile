FROM golang:1.23.4-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.AppVersion=$(cat VERSION)" -o /app/ogbsteam .

FROM alpine:3.18

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/ogbsteam /app/ogbsteam
COPY steam-config.yaml /app/steam-config.yaml

RUN chmod +x /app/ogbsteam && \
    chown -R appuser:appgroup /app

USER appuser

EXPOSE 12131

CMD ["/app/ogbsteam", "serve", "--config", "/app/steam-config.yaml"]

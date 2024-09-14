FROM golang:1.23-alpine AS builder
RUN apk add --no-cache build-base
WORKDIR /app
COPY . .
RUN CGO_ENABLED=1 GOOS=`go env GOHOSTOS` GOARCH=`go env GOHOSTARCH` go build -o out/bot -ldflags="-w -s" .

FROM alpine:latest
COPY --from=builder /app/out/bot /app/bot
COPY --from=builder /app/.env /app/.env
WORKDIR /app
EXPOSE 4853
CMD ["bot"]
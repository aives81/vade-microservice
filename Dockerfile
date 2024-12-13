FROM golang:1.23.4-alpine3.21 as builder

WORKDIR /app

COPY . .

RUN go build -o vade-doc-mngt-api

# Image finale
FROM alpine:3.13

COPY --from=builder /app/vade-doc-mngt-api /app/vade-doc-mngt-api

WORKDIR /app

EXPOSE 8001

# Point d'entrée
ENTRYPOINT ["./vade-doc-mngt-api"]

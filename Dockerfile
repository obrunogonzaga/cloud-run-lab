FROM golang:1.22 AS builder
LABEL authors="bruno"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY .env ./

WORKDIR /app/cmd

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/aplicacao .

FROM alpine:latest

EXPOSE 8000

# Instalar certificados CA
RUN apk --no-cache add ca-certificates

COPY --from=builder /bin/aplicacao /bin/aplicacao

COPY --from=builder /app/.env /

ENTRYPOINT ["/bin/aplicacao"]
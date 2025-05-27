FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ollama-proxy ./main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/ollama-proxy .
COPY --from=builder /app/config.yaml .

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 11434

ENTRYPOINT ["./ollama-proxy"]

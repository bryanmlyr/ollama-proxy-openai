FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o ollama-proxy ./main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/ollama-proxy .
COPY --from=builder /app/config.yaml .

EXPOSE 11434

ENTRYPOINT ["./ollama-proxy"]

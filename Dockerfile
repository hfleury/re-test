FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server cmd/main.go

# ---------- FINAL IMAGE ----------
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/config/configuration.yaml config/configuration.yaml

EXPOSE 8081

CMD ["./server"]

FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/backend ./cmd/dieton-backend/main.go

FROM ubuntu:latest AS runner
WORKDIR /app

RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*
COPY --from=builder /app/backend dieton-backend

EXPOSE 8080
CMD ["./dieton-backend"]
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/backend ./cmd/main.go

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/backend dieton-backend

EXPOSE 8080
CMD ["./dieton-backend"]
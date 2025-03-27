FROM golang:1.24.1-alpine AS builder
WORKDIR /app
COPY go/* ./
RUN go mod download
RUN go build -ldflags="-s -w" -o main .

FROM alpine:latest
EXPOSE 18080
WORKDIR /app
COPY --from=builder /app/main .
COPY go/.env .
CMD ["./main"]

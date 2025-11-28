# Etapa 1 — builder
FROM golang:1.25.4 as builder

WORKDIR /app

# Add go.mod and go.sum files
COPY go.mod ./ 
# RUN go mod download -- Use this only when go.sum was created

COPY . .

RUN go build -o server ./cmd/api

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]

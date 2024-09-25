# Build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/kisumu cmd/kisumu/main.go

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/kisumu /app/kisumu
LABEL Name=kisumu Version=0.0.1
EXPOSE 3000
ENTRYPOINT ["/app/kisumu"]
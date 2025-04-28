# FROM golang:1.23 AS builder
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod tidy
# COPY . .
# RUN CGO_ENABLED=0 go build -o auth-file-analyzer main.go
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# WORKDIR /root/
# COPY .env .  
# COPY --from=builder /app/auth-file-analyzer.
# EXPOSE 8002
# CMD ["./auth-file-analyzer"]

# Builder stage
FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 go build -o auth-file-analyzer main.go

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY .env .
COPY --from=builder /app/auth-file-analyzer .
EXPOSE 8002
CMD ["./auth-file-analyzer"]
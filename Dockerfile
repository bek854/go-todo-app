FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o todo-app .
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/todo-app .
CMD ["./todo-app"]

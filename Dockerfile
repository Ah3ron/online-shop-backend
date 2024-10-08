FROM golang:1.23.1 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -o main .

FROM alpine:latest
COPY --from=builder /app/main .
EXPOSE 3000
CMD ["./main"]


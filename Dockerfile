FROM golang:1.22.5
WORKDIR /app
COPY . .

RUN go mod tidy

RUN go build -o like ./cmd/like/main.go
CMD ["./like"]
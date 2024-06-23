FROM golang:1.22
LABEL authors="yar"

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod tidy

COPY . .

RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]

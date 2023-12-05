FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
COPY .env ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/http

EXPOSE 8080

CMD ["./main"]

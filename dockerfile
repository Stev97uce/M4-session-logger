FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

COPY .env .env

RUN go build -o session-logger .

EXPOSE 8000

CMD ["./session-logger"]

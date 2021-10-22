FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build .

CMD ["./pp-week2"]


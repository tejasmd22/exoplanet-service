FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ./configs /app/configs

COPY . .

RUN go build -o main main.go

EXPOSE 8010

CMD ["./main"]
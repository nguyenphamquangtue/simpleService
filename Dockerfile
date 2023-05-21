FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

# RUN go mod download

COPY ./cmd ./cmd

RUN go build -o main ./cmd

CMD ["./main"]

EXPOSE 3000

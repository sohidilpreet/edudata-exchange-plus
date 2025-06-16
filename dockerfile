FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN apt-get update && apt-get install -y libxml2-utils


RUN go build -o main ./cmd/*.go

CMD ["./main"]
